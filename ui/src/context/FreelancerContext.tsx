/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import React, { createContext, useCallback, useContext, useMemo, useRef, useState, useEffect } from 'react';
import request from '@/utils/request';

type PresenceMap = Record<string, boolean>;

interface FreelancerContextValue {
  hasProfile: (userId: string | undefined) => boolean;
  ensureProfileLoaded: (userId: string | undefined) => void;
}

const FreelancerContext = createContext<FreelancerContextValue | null>(null);

export const FreelancerProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [presence, setPresence] = useState<PresenceMap>({});
  const presenceRef = useRef<PresenceMap>({});
  const inFlight = useRef<Set<string>>(new Set());

  useEffect(() => {
    presenceRef.current = presence;
  }, [presence]);

  const hasProfile = useCallback(
    (userId: string | undefined) => {
      if (!userId) {
        return false;
      }
      return Boolean(presence[userId]);
    },
    [presence],
  );

  const ensureProfileLoaded = useCallback((userId: string | undefined) => {
    if (!userId) {
      return;
    }
    if (presenceRef.current[userId] !== undefined) {
      return;
    }
    if (inFlight.current.has(userId)) {
      return;
    }
    inFlight.current.add(userId);
    request
      .get('/answer/api/v1/freelancer/profile', { params: { user_id: userId } })
      .then(() => {
        setPresence((prev) => ({ ...prev, [userId]: true }));
      })
      .catch(() => {
        setPresence((prev) => ({ ...prev, [userId]: false }));
      })
      .finally(() => {
        inFlight.current.delete(userId);
      });
  }, []);

  const value = useMemo(() => ({ hasProfile, ensureProfileLoaded }), [hasProfile, ensureProfileLoaded]);

  return <FreelancerContext.Provider value={value}>{children}</FreelancerContext.Provider>;
};

export function useFreelancerContext() {
  const ctx = useContext(FreelancerContext);
  if (!ctx) {
    throw new Error('useFreelancerContext must be used within FreelancerProvider');
  }
  return ctx;
}

export function useHasFreelancerProfile(userId: string | undefined) {
  const { hasProfile, ensureProfileLoaded } = useFreelancerContext();
  useEffect(() => {
    ensureProfileLoaded(userId);
  }, [userId, ensureProfileLoaded]);
  return hasProfile(userId);
}


