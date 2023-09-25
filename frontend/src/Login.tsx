import { observer } from "mobx-react-lite";
import store from "./store/store";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import { Avatar, Box, Button, Container, CssBaseline, Link, TextField, Typography } from "@mui/material";
import SavingsOutlinedIcon from '@mui/icons-material/SavingsOutlined';
import FullScreenDialog from "./SignUp";
import Divider from '@mui/material/Divider';

function Login() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    let navigate = useNavigate();
    const login = async (event: React.FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault();
        if (username && password) {
            console.log("Вывод из формы: username:",username," password: ",password)
            await store.userStore.login(username, password);
            console.log("идем в main page")
            navigate("/", { replace: true });
        }
    };

    return (
        <Container className="login" component="main" maxWidth="xs">
            <CssBaseline />
            <Box
                sx={{
                    marginTop: 25,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                    <SavingsOutlinedIcon fontSize="medium" />
                </Avatar>
                <Typography component="h1" variant="h5">
                    Минибанк
                </Typography>
                <Box component="form" onSubmit={login} noValidate sx={{ mt: 1 }}>
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
                    <TextField size="small"
                        disabled
                        id="outlined-disabled"
                        label={store.userStore.authError}
                        defaultValue={store.userStore.authError}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                    >
                        Войти
                    </Button>
                    
                </Box>
                <Divider />
                <Box component="form" onSubmit={(e) => { e.preventDefault(); }} noValidate sx={{ mt: 1 }}>
                    <FullScreenDialog />
                </Box>
            </Box>
        </Container>
    );
}

const LoginObserver = observer(Login);
export default LoginObserver;
