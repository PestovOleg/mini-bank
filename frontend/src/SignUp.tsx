import * as React from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import CloseIcon from '@mui/icons-material/Close';
import Slide from '@mui/material/Slide';
import { TransitionProps } from '@mui/material/transitions';
import { Box, TextField } from '@mui/material';
import store from "./store/store";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import AlertSuccess from './AlertSignUp';

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement;
    },
    ref: React.Ref<unknown>,
) {
    return <Slide direction="up" ref={ref} {...props} />;
});

export default function FullScreenDialog() {
    const [open, setOpen] = React.useState(false);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [lastName, setLastName] = useState("");
    const [firstName, setFirstName] = useState("");
    const [patronymic, setPatronymic] = useState("");
    const [email, setEmail] = useState("");
    const [showAlert, setShowAlert] = React.useState(false);

    let navigate = useNavigate();
    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const signup = async (event: React.FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault();
        if (firstName && lastName && patronymic && email && username && password) {
            await store.userStore.signup(firstName, lastName, patronymic, email, username, password);
            
            // Show the alert
            setShowAlert(true);
            
            setTimeout(() => {
                navigate("/", { replace: true });
                setShowAlert(false);
            }, 5000);  
        }
    };
    
    return (
        <Box sx={{
            display: 'flex',
            flexDirection: 'column',
        }}
        >
            <Button variant="outlined" onClick={handleClickOpen}>
                Регистрация
            </Button>

            <Dialog
                open={open}
                onClose={handleClose}
                TransitionComponent={Transition}
            >
                <AppBar sx={{ position: 'relative',pr:2, pl:2 }}>
                    <Toolbar>
                        <IconButton
                            edge="start"
                            color="inherit"
                            onClick={handleClose}
                            aria-label="close"
                        >
                            <CloseIcon />
                        </IconButton>
                        <Typography sx={{ ml: 2, flex: 1 }} variant="h6" component="div">
                            Ввод данных
                        </Typography>
                        <Button autoFocus color="inherit" onClick={handleClose} >
                            ОК
                        </Button>
                    </Toolbar>
                </AppBar>
                <Box component="form" onSubmit={signup} noValidate sx={{ mt: 1,pr:5, pl:5 }}>
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="username"
                        label="Логин"
                        name="username"
                        value={username}
                        autoComplete="username"
                        autoFocus
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        name="password"
                        label="Password"
                        type="password"
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        autoComplete="current-password"
                    />
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="email"
                        label="email"
                        name="email"
                        value={email}
                        autoComplete="email"
                        autoFocus
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="firstName"
                        label="Имя"
                        name="firstName"
                        value={firstName}
                        autoComplete="firstName"
                        autoFocus
                        onChange={(e) => setFirstName(e.target.value)}
                    />
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="lastName"
                        label="Фамилия"
                        name="lastName"
                        value={lastName}
                        autoComplete="lastName"
                        autoFocus
                        onChange={(e) => setLastName(e.target.value)}
                    />
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="patronymic"
                        label="Отчество"
                        name="patronymic"
                        value={patronymic}
                        autoComplete="patronymic"
                        autoFocus
                        onChange={(e) => setPatronymic(e.target.value)}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                    >
                        Зарегистрироваться
                    </Button>
                    { showAlert && <AlertSuccess /> }
                </Box>
            </Dialog>
        </Box>
    );
}