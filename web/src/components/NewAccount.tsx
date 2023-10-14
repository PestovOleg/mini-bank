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
import { Alert, Box, Container, Link, Snackbar, TextField } from "@mui/material";
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

interface NewAccountDialogProps {
    open: boolean;
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  }

export default function NewAccountDialog({ open, setOpen }:NewAccountDialogProps) {
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
        setAccountName("");
        setCurrency("");  
    };

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
            setAccountName("");
            setCurrency("");          
            navigate("/", { replace: true });
            setShowAlert(true);
            setTimeout(() => {
                setShowAlert(false);
                handleClose();                
            }, 2000);
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
                            Открыть счет
                        </Typography>
                    </Toolbar>
                </AppBar>
                <Container className="mainPage" component="main" maxWidth="xs">
                <Box
                    component="form"
                    onSubmit={openAccount}
                    noValidate
                    sx={{ mt: 1}}
                >
                    
                    <Select
                        id="demo-simple-select"
                        value={currency}
                        label="Without label"
                        onChange={handleChange}
                        fullWidth
                        required                        
                        displayEmpty
                    >
                        <MenuItem disabled value="">
                            <em>Валюта счета </em>
                        </MenuItem>
                        <MenuItem key={810} value={"810"}>Рубль</MenuItem>
                        <MenuItem key={840} value={"840"}>Доллар</MenuItem>
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
                        disabled={!currency || !accountName}
                        sx={{ mt: 3, mb: 2 }}
                    >
                        Открыть счет
                    </Button>
                </Box>
                </Container>
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