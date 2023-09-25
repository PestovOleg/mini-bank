import * as React from 'react';
import Alert from '@mui/material/Alert';
import Stack from '@mui/material/Stack';

export default function AlertSuccess() {
  return (
    <Stack sx={{ width: '100%' }} spacing={2}>
      <Alert severity="info">Регситрация успешно! Выполните вход.</Alert>
    </Stack>
  );
}