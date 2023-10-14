import { useEffect, useState } from "react";
import store from "./store/store";
import { observer } from "mobx-react-lite";
import {
  AppBar,
  Avatar,
  BottomNavigationAction,
  Box,
  Card,
  CardContent,
  CardHeader,
  Container,
  CssBaseline,
  Divider,
  IconButton,
  Toolbar,
  Tooltip,
  Typography,
} from "@mui/material";
import SavingsOutlinedIcon from "@mui/icons-material/SavingsOutlined";
import { blue } from "@mui/material/colors";
import PhoneIcon from "@mui/icons-material/Phone";
import EmailIcon from "@mui/icons-material/Email";
import Account from "./components/Account";
import ModeSharpIcon from "@mui/icons-material/ModeSharp";
import GitHubIcon from "@mui/icons-material/GitHub";
import SpeedDialMenu from "./components/SpeedMenu";
import ChangeUserDetailsDialog from "./components/ChangeUserDetails";
import React from "react";
import LogoutIcon from '@mui/icons-material/Logout';
import DeleteUserDialog from "./components/DeleteUser";

function MainPage() {
  const [isLoading, setIsLoading] = useState(true);
  const [changeUserDetailsDialogOpen, setChangeUserDetailsDialogOpen] = React.useState(false);

  const openChangeUserDetailsDialog = () => {
    setChangeUserDetailsDialogOpen(true);
  };

  useEffect(() => {
    store.userStore.getUser().finally(() => setIsLoading(false));
    store.accountStore.getList().finally(() => setIsLoading(false));
  }, []);

  const user = store.userStore.User;
  const accountItems = store.accountStore.Accounts;
  const currentDate = () => new Date();

  function formatPhone(phoneNumber: string) {
    return phoneNumber.replace(
      /(\d{1})(\d{3})(\d{3})(\d{2})(\d{2})/,
      "+$1 ($2) $3-$4-$5"
    );
  }

  const accountList = () => {
    if ((store.toggleStore.getFeature("ListAccountsToggle")) && (accountItems.length>0)) {
      return (
        <Box sx={{ width: "100%" }}>
          <Account title="" accounts={accountItems} />
        </Box>
      );
    } else {
      return null;
    }
  };

  const userDetails = () => {
    if (store.toggleStore.getFeature("GetUserToggle")) {
      return (
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
      );
    } else {
      return null;
    }
  };

  const editUserDetails = () => {
    if (store.toggleStore.getFeature("UpdateUserToggle")) {
      return (
        <Tooltip title="Изменить контактные данные">
          <IconButton aria-label="settings" sx={{ m: 1 }} onClick={() => { openChangeUserDetailsDialog() }}>
            <ModeSharpIcon />
          </IconButton>
        </Tooltip>
      );
    } else {
      return null;
    }
  };

  const deleteUser =()=>{
    if (store.toggleStore.getFeature("DeleteUserToggle")) {
      return (
        <DeleteUserDialog />
      );
    } else {
      return null;
    }
  };
  
  if (isLoading) {
    return <div>Loading...</div>;
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
            <Tooltip title="Выход">
              <IconButton
                size="large"
                color="inherit"
                aria-label="menu"
                sx={{ ml: 5 }}
                onClick={() => { store.userStore.logout() }}
              >
                <LogoutIcon fontSize="large" />
              </IconButton>
            </Tooltip>
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

              <Box>{editUserDetails()}
          {deleteUser()}</Box>
              

            }
            title={
              user && user.name && user.lastName
                ? `${user.lastName} ${user.name}`
                : "Loading..."
            }
            titleTypographyProps={{ variant: "h6" }}
            subheader={currentDate().toLocaleDateString()}
          />
          {userDetails()}
        </Card>

        {accountList()}

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

              <Box sx={{ position: 'absolute' }}>
                <Box>
                  <IconButton color="inherit" aria-label="open drawer">
                    <BottomNavigationAction
                      component="a"
                      href="https://github.com/PestovOleg/mini-bank"
                      label="Github"
                      icon={<GitHubIcon />}
                    />
                  </IconButton>
                  <Typography variant="caption" sx={{ fontSize: 12, ml: -3 }}>
                    &copy; by Pestov Oleg
                  </Typography>
                </Box>
              </Box>
            </AppBar>

          </Container>
        </Box>

      </Box>
      <ChangeUserDetailsDialog
        open={changeUserDetailsDialogOpen}
        setOpen={setChangeUserDetailsDialogOpen}
      />
    </Container>
  );
}

const MainPageObserver = observer(MainPage);

export default MainPageObserver;
