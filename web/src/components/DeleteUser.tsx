import * as React from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { Alert, IconButton, Snackbar, Tooltip } from '@mui/material';
import PersonRemoveIcon from '@mui/icons-material/PersonRemove';
import store from '../store/store';
import { useNavigate } from 'react-router-dom';

export default function DeleteUserDialog() {
    const [open, setOpen] = React.useState(false);
    const [showAlert, setShowAlert] = React.useState(false);
    let navigate = useNavigate();

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    

    const deleteUser = async (
    ): Promise<void> => {
        //preventDefault();
            try {
                await store.userStore.deleteUser();
                // Показываем уведомление об успехе
                await store.userStore.logout();

                navigate("/login", { replace: true });
                setShowAlert(true);

                //setEmail("");
                //setPhone("");    
                setTimeout(() => {
                    setShowAlert(false);
                    handleClose();
                }, 2000);

            } catch (error) {
                // Выводим ошибку в консоль
                console.error("An error occurred:", error);
            }
        
    };

    return (
        <div>
            <Tooltip title="Удалить профиль">
                <IconButton
                    aria-label="settings"
                    sx={{ m: 1 }}
                    onClick={() => { handleClickOpen() }}
                >
                    <PersonRemoveIcon />
                </IconButton>
            </Tooltip>
            <Dialog
                open={open}
                onClose={handleClose}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                <DialogTitle id="alert-dialog-title">
                    {"Удалить профиль?"}
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        Профиль будет удален. Все счета будут удалены.
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}>Несогласен</Button>
                    <Button  onClick={() => { deleteUser() }} autoFocus>
                        Согласен
                    </Button>
                </DialogActions>
                {showAlert && (
                    <Snackbar open={open} onClose={handleClose}>
                        <Alert onClose={handleClose} severity="info" sx={{ width: "100%" }}>
                            Имя счета изменено успешно.
                        </Alert>
                    </Snackbar>
                )}
            </Dialog>
        </div>
    );
}

function preventDefault() {
    throw new Error('Function not implemented.');
}
