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
import InputMask from "react-input-mask";
import { formatDate } from "../utils/utils"
import { DateField } from "@mui/x-date-pickers/DateField";
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import dayjs from "dayjs";
import { deDE } from '@mui/x-date-pickers/locales';


const Transition = React.forwardRef(function Transition(
  props: TransitionProps & {
    children: React.ReactElement;
  },
  ref: React.Ref<unknown>
) {
  return <Slide direction="up" ref={ref} {...props} />;
});

export default function SignUpDialog() {
  const [open, setOpen] = React.useState(false);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [lastName, setLastName] = useState("");
  const [firstName, setFirstName] = useState("");
  const [patronymic, setPatronymic] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [showAlert, setShowAlert] = React.useState(false);
  const [birthday, setBirthday] = useState<dayjs.Dayjs | null>(null);


  let navigate = useNavigate();
  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const signup = async (
    event: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    event.preventDefault();
    const formattedBirthday = birthday ? birthday.format('DD.MM.YYYY') : null;
    if (
      firstName &&
      lastName &&
      patronymic &&
      email &&
      username &&
      password &&
      phone &&
      formattedBirthday
    ) {
      await store.userStore.signup(
        firstName,
        lastName,
        patronymic,
        email,
        username,
        password,
        phone,
        formattedBirthday
      );

      // Show the alert
      setShowAlert(true);

      setTimeout(() => {
        navigate("/", { replace: true });
        setShowAlert(false);
      }, 5000);
    }
  };


  const handleEmailChange = (value: string) => {
    // Remove all characters except digits, dots, lowercase, uppercase, and @
    value = value.replace(/[^0-9.^a-z^A-Z@]/g, "");

    setEmail(value);
  };

  const isEmailValid = (email: string) => {
    if (email === "") return true;  // Don't flag an empty string as invalid
  
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    return emailRegex.test(email);
  };
  

  const handleLoginChange = (value: string) => {
    value = value.replace(/[^0-9^a-z^A-Z]/g, "");
    setUsername(value);
  };

  return (

    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
      }}
    >
      <Link component="button" variant="body2" onClick={handleClickOpen}>
        Регистрация
      </Link>

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
              Регистрация
            </Typography>
          </Toolbar>
        </AppBar>
        <Container className="mainPage" component="main" maxWidth="xs">
          <Box
            component="form"
            onSubmit={signup}
            noValidate
          >
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
            <TextField
              margin="normal"
              required
              fullWidth
              id="firstName"
              label="Имя"
              name="firstName"
              value={firstName}
              autoComplete="firstName"
              autoFocus
              onChange={(e) => setFirstName(e.target.value)}
            />
            <TextField
              margin="normal"
              required
              fullWidth
              id="lastName"
              label="Фамилия"
              name="lastName"
              value={lastName}
              autoComplete="lastName"
              autoFocus
              onChange={(e) => setLastName(e.target.value)}
            />
            <TextField
              margin="normal"
              required
              fullWidth
              id="patronymic"
              label="Отчество"
              name="patronymic"
              value={patronymic}
              autoComplete="patronymic"
              autoFocus
              onChange={(e) => setPatronymic(e.target.value)}
            />
            <TextField
              margin="normal"
              required
              fullWidth
              error={!isEmailValid(email)}
              helperText={!isEmailValid(email) ? 'Invalid email!' : ' '}
              id="email"
              label="email"
              name="email"
              value={email}
              autoComplete="email"
              autoFocus
              onChange={(e) => handleEmailChange(e.target.value)}
            />
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
            <LocalizationProvider dateAdapter={AdapterDayjs} adapterLocale='de'>
              <DateField
                format="DD.MM.YYYY"
                onChange={(date) => setBirthday(date)}
                fullWidth
                required
                defaultValue={dayjs('2022-04-17')}
              />
            </LocalizationProvider>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Зарегистрироваться
            </Button>
          </Box>
        </Container>
        {showAlert && (
          <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="info" sx={{ width: "100%" }}>
              Регистрация завершена, выполните вход.
            </Alert>
          </Snackbar>

        )}


      </Dialog>
    </Box>

  );
}