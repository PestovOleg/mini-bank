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
import { Alert, Box, Container, Snackbar, TextField } from "@mui/material";
import store from "../store/store";
import { useNavigate } from "react-router-dom";
import InputLabel from '@mui/material/InputLabel';
import { IAccount } from "../models/types";

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement;
    },
    ref: React.Ref<unknown>
) {
    return <Slide direction="up" ref={ref} {...props} />;
});

interface ChangeAccountNameDialogProps {
    open: boolean;
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
    account: IAccount | null;
}


export default function ChangeAccountNameDialog({ open, setOpen, account }: ChangeAccountNameDialogProps ) {
    const [showAlert, setShowAlert] = React.useState(false);
    const [accountName, setAccountName] = React.useState("");
    let navigate = useNavigate();

    const handleClose = () => {
        setOpen(false);
    };


    const changeAccountName = async (
        event: React.FormEvent<HTMLFormElement>
    ): Promise<void> => {
        event.preventDefault();
        if (account!=null) {
            try {
                await store.accountStore.updateAccount(
                    accountName,
                    String(account.id)
                );
                // Показываем уведомление об успехе

                setAccountName("");
                navigate("/", { replace: true });
                setShowAlert(true);
                store.accountStore.getList();

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
                            Сменить имя счета
                        </Typography>
                    </Toolbar>
                </AppBar>
                <Container className="mainPage" component="main" maxWidth="xs">
                <Box
                    component="form"
                    onSubmit={changeAccountName}
                    noValidate
                    sx={{ mt: 1}}
                >
                    <InputLabel htmlFor="accountName"></InputLabel>

                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="accountName"
                        label="Название счета"
                        name="accountName"
                        value={account ? account.name : ""}
                        disabled
                    />
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="accountName"
                        label="Название счета"
                        name="accountName"
                        value={accountName}
                        autoFocus
                        onChange={(e) => setAccountName(e.target.value)}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                        disabled={!accountName}
                    >
                        Изменить
                    </Button>
                </Box>
                </Container>
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