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

interface NewPaymentDialogProps {
    open: boolean;
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

export default function TopupDialog({ open, setOpen }: NewPaymentDialogProps) {
    const [showAlert, setShowAlert] = React.useState(false);
    const [selectedAccountFrom, setSelectedAccountFrom] = React.useState<IAccount | null>(null);
    const [amount, setAmount] = React.useState("");

    let navigate = useNavigate();

    const handleClose = () => {
        setOpen(false);
    };

    const checkAmountForTopup = () => {
        if (Number(amount) > 0) return false;
    };

    const handleAmountChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        let value = e.target.value;

        // Удалить все символы, кроме цифр и точки
        value = value.replace(/[^0-9.]/g, "");

        // Проверка на наличие более чем одной точки
        const dotCount = value.split(".").length - 1;
        if (dotCount > 1) {
            value = value.slice(0, value.lastIndexOf("."));
        }

        // Ограничить количество знаков после запятой до двух
        if (dotCount === 1) {
            const [integerPart, decimalPart] = value.split(".");
            if (decimalPart.length > 2) {
                value = `${integerPart}.${decimalPart.slice(0, 2)}`;
            }
        }

        setAmount(value);
    };

    const accountItems = React.useMemo(() => {
        return store.accountStore.Accounts.sort((a, b) => b.amount - a.amount);
      }, [store.accountStore.Accounts]);

    const makePayment = async (
        event: React.FormEvent<HTMLFormElement>
    ): Promise<void> => {
        event.preventDefault();
        if (selectedAccountFrom) {
            if (selectedAccountFrom.account && Number(amount)) {
                try {
                    await store.accountStore.topUpAccount(
                        Number(amount),
                        String(selectedAccountFrom.id)
                    );
    
                    // Показываем уведомление об успехе
                    
                    setSelectedAccountFrom(null);
                    setAmount("");
                    
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
                            Пополнить счет
                        </Typography>
                    </Toolbar>
                </AppBar>
                <Box
                    component="form"
                    onSubmit={makePayment}
                    noValidate
                    sx={{ mt: 1, pr: 5, pl: 5 }}
                >
                    <AccountSelect
                        accounts={accountItems}
                        placeHolder="Счет"
                        style={{ margin: '50px' }}
                        onAccountSelected={(account) => setSelectedAccountFrom(account)}
                    ></AccountSelect>
                    <InputLabel htmlFor="outlined-adornment-amount"></InputLabel>

                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="amount"
                        label="Сумма"
                        value={amount}
                        onChange={handleAmountChange}
                    />

                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                        disabled={!selectedAccountFrom || checkAmountForTopup()}
                    >
                        Пополнить
                    </Button>
                </Box>
                {showAlert && (
                    <Snackbar open={open} onClose={handleClose}>
                        <Alert onClose={handleClose} severity="info" sx={{ width: "100%" }}>
                            Пополнение баланса осуществлено успешно.
                        </Alert>
                    </Snackbar>
                )}
            </Dialog>
        </Box>
    );
}