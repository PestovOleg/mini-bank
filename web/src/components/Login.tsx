import { observer } from "mobx-react-lite";
import store from "../store/store";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import { AppBar, Avatar, Box, Button, Container, CssBaseline, FormHelperText, IconButton, Link, TextField, Toolbar, Typography } from "@mui/material";
import SavingsOutlinedIcon from '@mui/icons-material/SavingsOutlined';
import SignUpDialog from "./SignUp";

function Login() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    let navigate = useNavigate();
    const login = async (event: React.FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault();
        if (username && password) {
            await store.userStore.login(username, password);
            navigate("/", { replace: true });

        }
    };

    const signUp = () => {
        if (store.toggleStore.getFeature("CreateUserToggle")) {
            return (
                <Box component="form" onSubmit={(e) => { e.preventDefault(); }} noValidate sx={{ mt: 1 }}>
                    <SignUpDialog />
                </Box>
            );
        } else {
            return null;
        }
    };

    const handleLoginChange = (value: string) => {
        value = value.replace(/[^0-9^a-z^A-Z]/g, "");
        setUsername(value);
    };
    
    return (
        <Container className="login" component="main" maxWidth="xs" >
            <CssBaseline />
            <Box
                sx={{
                    marginTop: 25,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    height: '100%',
                }}
            >
                <AppBar position="static">
                    <Toolbar>
                        <IconButton
                            size="large"
                            edge="start"
                            color="inherit"
                            aria-label="menu"
                            sx={{ mr: 2 }}
                        >
                            <SavingsOutlinedIcon fontSize="large" />
                        </IconButton>
                        <Typography variant="h6" component="div" sx={{ ml: 5, letterSpacing: 5 }}>
                            МИНИБАНК
                        </Typography>

                    </Toolbar>
                </AppBar>
                
                <Box component="form" onSubmit={login} noValidate sx={{ mt: 1}}>
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
                        helperText="Только цифры и латинские символы"
                        onChange={(e) => handleLoginChange(e.target.value)}
                    />
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        name="password"
                        label="Пароль"
                        type="password"
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        autoComplete="current-password"
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        size="large"
                        sx={{ mt: 3, mb: 2, }}
                    >
                        Войти
                    </Button>

                </Box>
               
                {signUp()}
            </Box>
        </Container>
    );
}

const LoginObserver = observer(Login);
export default LoginObserver;
