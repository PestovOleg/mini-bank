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
import InputLabel from '@mui/material/InputLabel';
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

interface CloseAccountDialogProps {
    open: boolean;
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
    account: IAccount | null;
}

export default function CloseAccountDialog({ open, setOpen,account }: CloseAccountDialogProps) {
    const [showAlert, setShowAlert] = React.useState(false);
    const [selectedAccountTo, setSelectedAccountTo] = React.useState<IAccount | null>(null);

    let navigate = useNavigate();

    const handleClose = () => {
        setOpen(false);
    };    

    const accountItems = React.useMemo(() => {
        return store.accountStore.Accounts.sort((a, b) => b.amount - a.amount);
      }, [store.accountStore.Accounts]);

    const makePayment = async (
        event: React.FormEvent<HTMLFormElement>
    ): Promise<void> => {
        event.preventDefault();       
            if (account!=null && selectedAccountTo!=null) {
                try {
                    await store.accountStore.closeAccount(
                        String(account.id)
                    );
                    
                    if (Number(account.amount>0)){
                        await store.accountStore.topUpAccount(
                            Number(account.amount),
                            String(selectedAccountTo.id)
    
                        );
                    }
                    
                    // Показываем уведомление об успехе
                                        
                    navigate("/", { replace: true });  
                    setShowAlert(true);              
                    
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
                            Закрыть счет
                        </Typography>
                    </Toolbar>
                </AppBar>
                <Box
                    component="form"
                    onSubmit={makePayment}
                    noValidate
                    sx={{ mt: 1, pr: 5, pl: 5 }}           >    
                    <InputLabel htmlFor="accountName">Текущее имя</InputLabel>
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="accountName"
                        label="Закрыть счет"
                        name="accountName"
                        value={account ? account.name : ""}
                        disabled
                    />
                     <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="accountName"
                        label="Будет списана сумма"
                        name="accountName"
                        value={account ? account.amount : ""}
                        disabled
                    />
                    <AccountSelect
                        accounts={accountItems}
                        placeHolder="Зачислить остаток на счет"
                        style={{ margin: '50px' }}
                        onAccountSelected={(account) => setSelectedAccountTo(account)}
                    ></AccountSelect>
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                        disabled={!selectedAccountTo}
                    >
                        Закрыть счет
                    </Button>
                </Box>
                {showAlert && (
                    <Snackbar open={open} onClose={handleClose}>
                        <Alert onClose={handleClose} severity="info" sx={{ width: "100%" }}>
                            Счет закрыт успешно.
                        </Alert>
                    </Snackbar>
                )}
            </Dialog>
        </Box>
    );
}