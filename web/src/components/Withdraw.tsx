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
import { Alert, Box,  Snackbar, TextField } from "@mui/material";
import store from "../store/store";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import RedeemIcon from '@mui/icons-material/Redeem';
import { styled } from "@mui/material/styles";
import Fab from "@mui/material/Fab";
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import Select, { SelectChangeEvent } from '@mui/material/Select';

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement;
    },
    ref: React.Ref<unknown>
) {
    return <Slide direction="up" ref={ref} {...props} />;
});

export default function WithdrawDialog() {
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
        left: 50,
        right: 0,
        margin: "0",
    });

    const openAccount = async (
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
            navigate("/", { replace: true });
            setTimeout(() => {
                
                setShowAlert(false);
            }, 1000);
        }
    };

    return (
        <Box
            sx={{
                display: "flex",
                flexDirection: "column",
            }}
        >
            <StyledFab color="info" aria-label="add" onClick={handleClickOpen}>
                <RedeemIcon />
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
                    onSubmit={openAccount}
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
                    <Snackbar open={open} onClose={handleClose}>
                        <Alert onClose={handleClose} severity="info" sx={{ width: "100%" }}>
                            Счет открыт.
                        </Alert>
                    </Snackbar>
                )}
            </Dialog>
        </Box>
    );
}