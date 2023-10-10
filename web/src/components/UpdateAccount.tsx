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
import { Alert, Box, Link, Snackbar, TextField } from "@mui/material";
import store from "../store/store";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import AlertSuccess from "./AlertSignUp";
import InputMask from "react-input-mask";
import useMediaQuery from "@mui/material/useMediaQuery";
import { formatDate } from "../utils/utils"
import StyledFab from "@mui/material/Fab";
import AddIcon from "@mui/icons-material/Add";
import { styled } from "@mui/material/styles";
import Fab from "@mui/material/Fab";
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement;
    },
    ref: React.Ref<unknown>
) {
    return <Slide direction="up" ref={ref} {...props} />;
});

export default function UpdateAccountDialog() {
    const [open, setOpen] = React.useState(false);
    const [accountName, setAccountName] = useState("");
    const [showAlert, setShowAlert] = React.useState(false);
    const [currency, setCurrency] = React.useState("");

    const handleChange = (event: SelectChangeEvent) => {
        setCurrency(event.target.value as string);
    };

    let navigate = useNavigate();
    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const StyledFab = styled(Fab)({
        position: "absolute",
        zIndex: 1,
        top: -30,
        left: 0,
        right: 0,
        margin: "0 auto",
    });

    const signup = async (
        event: React.FormEvent<HTMLFormElement>
    ): Promise<void> => {
        event.preventDefault();
        if (
            currency &&
            accountName
        ) {
            await store.accountStore.openAccount(
                currency,
                accountName
            );

            // Show the alert
            setShowAlert(true);

            setTimeout(() => {
                navigate("/", { replace: true });
                setShowAlert(false);
            }, 5000);
        }
    };

    return (
        <Box
            sx={{
                display: "flex",
                flexDirection: "column",
            }}
        >
            <StyledFab color="secondary" aria-label="add" onClick={handleClickOpen}>
                <AddIcon />
            </StyledFab>

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
                            Открыть счет
                        </Typography>
                    </Toolbar>
                </AppBar>
                <Box
                    component="form"
                    onSubmit={signup}
                    noValidate
                    sx={{ mt: 1, pr: 5, pl: 5 }}
                >
                    <InputLabel id="demo-simple-select-label" >Валюта</InputLabel>
                    <Select
                        labelId="demo-simple-select-label"
                        id="demo-simple-select"
                        value={currency}
                        label="Валюта"
                        onChange={handleChange}
                        autoWidth
                        required
                    >
                        <MenuItem value={"810"}>Рубль</MenuItem>
                        <MenuItem value={"840"}>Доллар</MenuItem>
                    </Select>
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
                    >
                        Открыть счет
                    </Button>
                </Box>
                {showAlert && (
                    <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                        <Alert onClose={handleClose} severity="info" sx={{ width: "100%" }}>
                            Регистрация завершена,выполните вход.
                        </Alert>
                    </Snackbar>
                )}
            </Dialog>
        </Box>
    );
}