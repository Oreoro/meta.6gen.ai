import React, { useState } from 'react';
import { Modal, Button, Form, Alert } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import request from '@/utils/request';

interface Props {
  freelancerUserId: string;
  freelancerDisplayName: string;
  className?: string;
}

const HireMeButton: React.FC<Props> = ({
  freelancerUserId,
  freelancerDisplayName,
  className = '',
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'freelancer' });
  const [showModal, setShowModal] = useState(false);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');
  const [subject, setSubject] = useState('');
  const [alert, setAlert] = useState<{
    type: 'success' | 'danger';
    message: string;
  } | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!subject.trim() || !message.trim()) {
      setAlert({ type: 'danger', message: t('form.required_fields') });
      return;
    }

    setLoading(true);
    setAlert(null);

    try {
      await request.post('/answer/api/v1/freelancer/hire', {
        freelancer_user_id: freelancerUserId,
        subject: subject.trim(),
        message: message.trim(),
      });

      setAlert({ type: 'success', message: t('hire.success_message') });
      setTimeout(() => {
        setShowModal(false);
        setSubject('');
        setMessage('');
        setAlert(null);
      }, 2000);
    } catch (error) {
      setAlert({ type: 'danger', message: t('hire.error_message') });
    } finally {
      setLoading(false);
    }
  };

  const handleClose = () => {
    setShowModal(false);
    setSubject('');
    setMessage('');
    setAlert(null);
  };

  return (
    <>
      <Button
        variant="success"
        size="sm"
        className={`hire-me-btn ${className}`}
        onClick={() => setShowModal(true)}>
        💼 {t('hire.button_text')}
      </Button>

      <Modal show={showModal} onHide={handleClose} size="lg">
        <Modal.Header closeButton>
          <Modal.Title>
            {t('hire.modal_title', { name: freelancerDisplayName })}
          </Modal.Title>
        </Modal.Header>
        <Modal.Body>
          {alert && (
            <Alert variant={alert.type} className="mb-3">
              {alert.message}
            </Alert>
          )}
          <Form onSubmit={handleSubmit}>
            <Form.Group className="mb-3">
              <Form.Label>{t('hire.subject_label')}</Form.Label>
              <Form.Control
                type="text"
                value={subject}
                onChange={(e) => setSubject(e.target.value)}
                placeholder={t('hire.subject_placeholder')}
                required
              />
            </Form.Group>

            <Form.Group className="mb-3">
              <Form.Label>{t('hire.message_label')}</Form.Label>
              <Form.Control
                as="textarea"
                rows={6}
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                placeholder={t('hire.message_placeholder', {
                  name: freelancerDisplayName,
                })}
                required
              />
            </Form.Group>

            <div className="d-flex justify-content-end gap-2">
              <Button
                variant="secondary"
                onClick={handleClose}
                disabled={loading}>
                {t('common.cancel')}
              </Button>
              <Button variant="success" type="submit" disabled={loading}>
                {loading ? t('common.sending') : t('hire.send_button')}
              </Button>
            </div>
          </Form>
        </Modal.Body>
      </Modal>
    </>
  );
};

export default HireMeButton;
