import * as React from "react";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import CloseIcon from "@mui/icons-material/Close";
import Slide from "@mui/material/Slide";
import { TransitionProps } from "@mui/material/transitions";
import { Alert, Box, Snackbar, TextField } from "@mui/material";
import store from "../store/store";
import { useNavigate } from "react-router-dom";
import InputMask from "react-input-mask";
import InputLabel from '@mui/material/InputLabel';

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement;
    },
    ref: React.Ref<unknown>
) {
    return <Slide direction="up" ref={ref} {...props} />;
});

interface ChangeUserDetailDialogProps {
    open: boolean;
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}


export default function ChangeUserDetailsDialog({ open, setOpen }: ChangeUserDetailDialogProps) {
    const [showAlert, setShowAlert] = React.useState(false);
    const [email, setEmail] = React.useState(store.userStore.User.email);
    const [phone, setPhone] = React.useState(store.userStore.User.phone);
    let navigate = useNavigate();

    const handleClose = () => {
        setOpen(false);
    };


    const changeUserDetails = async (
        event: React.FormEvent<HTMLFormElement>
    ): Promise<void> => {
        event.preventDefault();
        if (email !== "" && phone !== "") {
            try {
                await store.userStore.updateUser(
                    email,
                    phone
                );
                // Показываем уведомление об успехе
                await store.userStore.getUser();

                navigate("/", { replace: true });
                setShowAlert(true);

                //setEmail("");
                //setPhone("");    
                setTimeout(() => {
                    setShowAlert(false);
                    handleClose();
                }, 1000);

            } catch (error) {
                // Выводим ошибку в консоль
                console.error("An error occurred:", error);
            }
        }
    };


    return (
        <Box
            sx={{
                display: "flex",
                flexDirection: "column",
            }}
        >

            <Dialog
                open={open}
                onClose={handleClose}
                TransitionComponent={Transition}
            >
                <AppBar sx={{ position: "relative", pr: 2, pl: 2 }}>
                    <Toolbar>
                        <IconButton
                            edge="start"
                            color="inherit"
                            onClick={handleClose}
                            aria-label="close"
                        >
                            <CloseIcon />
                        </IconButton>
                        <Typography
                            sx={{ textAlign: "center", flex: 1, mr: 5 }}
                            variant="h6"
                            component="div"
                        >
                            Сменить контактные данные
                        </Typography>
                    </Toolbar>
                </AppBar>
                <Box
                    component="form"
                    onSubmit={changeUserDetails}
                    noValidate
                    sx={{ mt: 1, pr: 5, pl: 5 }}
                >
                    <InputLabel htmlFor="email">
                        
                    </InputLabel>

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

                    <InputLabel htmlFor="phone"></InputLabel>
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="phone"
                        label="Телефон"
                        name="phone"
                        value={phone}
                        autoComplete="phone"
                        autoFocus
                        onChange={(e) => setPhone(e.target.value)}
                        InputProps={{
                            inputComponent: InputMask,
                            inputProps: { mask: "+7 (999) 999-99-99" },
                        }}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                        disabled={!email && !phone}
                    >
                        Изменить
                    </Button>
                </Box>
                {showAlert && (
                    <Snackbar open={open} onClose={handleClose}>
                        <Alert onClose={handleClose} severity="info" sx={{ width: "100%" }}>
                            Имя счета изменено успешно.
                        </Alert>
                    </Snackbar>
                )}
            </Dialog>
        </Box>
    );
}