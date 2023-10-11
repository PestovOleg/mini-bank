import { useEffect, useState } from "react";
import store from "./store/store";
import { styled } from "@mui/material/styles";
import { observer } from "mobx-react-lite";
import {
  AppBar,
  Avatar,
  BottomNavigation,
  BottomNavigationAction,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  CardMedia,
  Container,
  CssBaseline,
  Divider,
  Grid,
  IconButton,
  Link,
  List,
  ListItem,
  ListItemAvatar,
  ListItemText,
  Paper,
  TextField,
  Toolbar,
  Typography,
} from "@mui/material";
import SavingsOutlinedIcon from "@mui/icons-material/SavingsOutlined";
import { blue } from "@mui/material/colors";
import PhoneIcon from "@mui/icons-material/Phone";
import EmailIcon from "@mui/icons-material/Email";
import Account from "./components/Account";
import ModeSharpIcon from "@mui/icons-material/ModeSharp";
import GitHubIcon from "@mui/icons-material/GitHub";
import Fab from "@mui/material/Fab";
import NewAccountDialog from "./components/NewAccount";
import PaymentDialog from "./components/Payment";
import SpeedDialMenu from "./components/SpeedMenu";

function MainPage() {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Сделайте функцию getUser асинхронной или вызовите другую асинхронную функцию
    store.userStore.getUser().finally(() => setIsLoading(false)); // измените эту строку
    store.accountStore.getList().finally(() => setIsLoading(false));
  }, []);

  const user = store.userStore.User;
  const accountItems = store.accountStore.Accounts;
  const currentDate = () => new Date();
  const StyledFab = styled(Fab)({
    position: "absolute",
    zIndex: 1,
    top: -30,
    left: 0,
    right: 0,
    margin: "0 auto",
  });

  function formatPhone(phoneNumber: string) {
    return phoneNumber.replace(
      /(\d{1})(\d{3})(\d{3})(\d{2})(\d{2})/,
      "+$1 ($2) $3-$4-$5"
    );
  }

  if (isLoading) {
    return <div>Loading...</div>; // или другой компонент "загрузка"
  }
  return (
    <Container className="mainPage" component="main" maxWidth="xs">
      <CssBaseline />
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <AppBar position="sticky">
          <Toolbar>
            <IconButton
              size="large"
              edge="start"
              color="inherit"
              aria-label="menu"
              sx={{ mr: 2 }}
            >
              <SavingsOutlinedIcon fontSize="large" />
            </IconButton>
            <Typography
              variant="h6"
              component="div"
              sx={{ ml: 5, letterSpacing: 5 }}
            >
              МИНИБАНК
            </Typography>
          </Toolbar>
        </AppBar>
        <Card sx={{ mt: 1, width: "100%" }}>
          <CardHeader
            avatar={
              <Avatar sx={{ m: 1, bgcolor: blue[500] }}>
                {user && user.name && user.lastName
                  ? `${user.name[0]}${user.lastName[0]}`
                  : "..."}
              </Avatar>
            }
            action={
              <IconButton aria-label="settings" sx={{ m: 1 }}>
                <ModeSharpIcon />
              </IconButton>
            }
            title={
              user && user.name && user.lastName
                ? `${user.lastName} ${user.name}`
                : "Loading..."
            }
            titleTypographyProps={{ variant: "h6" }}
            subheader={currentDate().toLocaleDateString()}
          />
          <CardContent sx={{ mt: -2, mb: -2, width: "100%" }}>
            <Divider></Divider>
            <Box display="flex" alignItems="center" sx={{ mr: 1 }}>
              <IconButton aria-label="settings" sx={{ ml: 1 }}>
                <PhoneIcon />
              </IconButton>
              <Typography
                variant="body2"
                color="text.secondary"
                sx={{ marginLeft: 3, color: "black" }}
              >
                {user && user.phone
                  ? formatPhone(user.phone)
                  : "No phone available"}
              </Typography>
            </Box>
            <Box display="flex" alignItems="center">
              <IconButton aria-label="settings" sx={{ ml: 1 }}>
                <EmailIcon />
              </IconButton>
              <Typography
                variant="body2"
                color="text.secondary"
                sx={{ marginLeft: 3, color: "black" }}
              >
                {user && user.email ? user.email : "No email available"}
              </Typography>
            </Box>
          </CardContent>
        </Card>

        <Box sx={{ width: "100%" }}>
          <Account title="" accounts={accountItems} />
        </Box>

        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            justifyContent: 'center', // Centers the AppBar horizontally
            position: 'fixed',
            bottom: 0,
            width: '100%'
          }}
        >
          <Container maxWidth="xs">
            <AppBar
              position="static"
              color="primary"
            >
              <Toolbar sx={{ justifyContent: 'center' }}>
                
                
              <SpeedDialMenu />
                <Box sx={{ flexGrow: 1 }} />
              </Toolbar>
              
              <Box sx={{ position: 'absolute',marginLeft:17,marginTop:1 }}>
                  <IconButton color="inherit" aria-label="open drawer">
                    <BottomNavigationAction
                      component="a"
                      href="https://github.com/PestovOleg/mini-bank"
                      label="Github"
                      icon={<GitHubIcon />}
                    />
                  </IconButton>
                  <Typography variant="caption" sx={{ fontSize: 12, ml: -3 }}>
                    &copy; by Pestov
                  </Typography>
                </Box>
            </AppBar>
            
          </Container>
        </Box>

      </Box>
    </Container>
  );
}

const MainPageObserver = observer(MainPage);

export default MainPageObserver;
