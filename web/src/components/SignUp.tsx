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
import { Alert, Box, Container,  Link, Snackbar, TextField } from "@mui/material";
import store from "../store/store";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import InputMask from "react-input-mask";
import { formatDate } from "../utils/utils"

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
  const [birthday, setBirthday] = useState("");
  const [showAlert, setShowAlert] = React.useState(false);

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
    const b = new Date(birthday);
    const formattedDate = formatDate(b);
    if (
      firstName &&
      lastName &&
      patronymic &&
      email &&
      username &&
      password &&
      phone &&
      birthday
    ) {
      await store.userStore.signup(
        firstName,
        lastName,
        patronymic,
        email,
        username,
        password,
        phone,
        formattedDate
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

    // Удалить все символы, кроме цифр и точки
    value = value.replace(/[^0-9.^a-z^A-Z@]/g, "");

    // Проверка на наличие более чем одной точки
    const dotCount = value.split(".").length - 1;
    if (dotCount > 1) {
      value = value.slice(0, value.lastIndexOf("."));
    }

    // Проверка на наличие более чем одной собаки
    const dogCount = value.split("@").length - 1;
    if (dogCount > 1) {
      value = value.slice(0, value.lastIndexOf("."));
    }

    setEmail(value);
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
            <TextField
              margin="normal"
              required
              fullWidth
              type="date"
              id="birthday"
              name="birthday"
              value={birthday}
              autoFocus
              onChange={(e) => setBirthday(e.target.value)}
            />

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
              Регистрация завершена,выполните вход.
            </Alert>
          </Snackbar>
        )}


      </Dialog>
    </Box>

  );
}