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
import { Alert, Box, InputAdornment, Link, OutlinedInput, Snackbar, TextField } from "@mui/material";
import store from "../store/store";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import AlertSuccess from "./AlertSignUp";
import InputMask from "react-input-mask";
import useMediaQuery from "@mui/material/useMediaQuery";
import { formatDate } from "../utils/utils"
import StyledFab from "@mui/material/Fab";
import PaymentIcon from '@mui/icons-material/Payment';
import { styled } from "@mui/material/styles";
import Fab from "@mui/material/Fab";
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import AccountSelect from "./AccountSelect";
import { IAccount } from "../models/types";

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement;
    },
    ref: React.Ref<unknown>
) {
    return <Slide direction="up" ref={ref} {...props} />;
});

interface NewPaymentDialogProps {
    open: boolean;
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

export default function PaymentDialog({ open, setOpen }: NewPaymentDialogProps) {
    const [accountName, setAccountName] = useState("");
    const [showAlert, setShowAlert] = React.useState(false);
    const [selectedAccountTo, setSelectedAccountTo] = React.useState<IAccount | null>(null);
    const [selectedAccountFrom, setSelectedAccountFrom] = React.useState<IAccount | null>(null);

    
    let navigate = useNavigate();
    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const accountItems = store.accountStore.Accounts;

    const openAccount = async (
        event: React.FormEvent<HTMLFormElement>
    ): Promise<void> => {
        event.preventDefault();
        if (selectedAccountFrom && selectedAccountTo){
            if (
                selectedAccountFrom.account &&
                selectedAccountTo.account
            ) {
                await store.accountStore.withdrawAccount(
                    selectedAccountFrom.amount,
                    selectedAccountFrom.account
                    
                );
        }
        

            // Show the alert
            setShowAlert(true);
            navigate("/", { replace: true });
            setTimeout(() => {

                setShowAlert(false);
            }, 3000);
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
                            Перевод между счетами
                        </Typography>
                    </Toolbar>
                </AppBar>
                <Box
                    component="form"
                    onSubmit={openAccount}
                    noValidate
                    sx={{ mt: 1, pr: 5, pl: 5 }}
                >
                    <AccountSelect
                        accounts={accountItems}
                        placeHolder="Со счета"
                        style={{ margin: '50px' }}
                        onAccountSelected={(account) => setSelectedAccountFrom(account)}
                    ></AccountSelect>
                    <AccountSelect
                        accounts={accountItems}
                        placeHolder="На счет"
                        onAccountSelected={(account) => setSelectedAccountTo(account)}
                    ></AccountSelect>
                    <InputLabel htmlFor="outlined-adornment-amount"></InputLabel>
                    <OutlinedInput
                        id="outlined-adornment-amount"
                        fullWidth
                        startAdornment={<InputAdornment position="start">$</InputAdornment>}
                        
                        
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                        disabled={!selectedAccountTo || !selectedAccountFrom}
                    >
                        Перевести
                    </Button>
                </Box>
                {showAlert && (
                    <Snackbar open={open} onClose={handleClose}>
                        <Alert onClose={handleClose} severity="info" sx={{ width: "100%" }}>
                            Перевод осуществлен.
                        </Alert>
                    </Snackbar>
                )}
            </Dialog>
        </Box>
    );
}