import { useEffect } from "react";
import store from "./store/store";
import { observer } from "mobx-react-lite";
import { Avatar, Box, Button, Container, CssBaseline, Link, TextField, Typography } from "@mui/material";
import SavingsOutlinedIcon from '@mui/icons-material/SavingsOutlined';
import { blue, deepPurple } from "@mui/material/colors";

function MainPage() {
  useEffect(() => {
    store.userStore.getUser()
  }, [])

  const user = store.userStore.User;
  return (
    <Container className="login" component="main" maxWidth="xs">
      <CssBaseline />
      <Box
        sx={{
          marginTop: 25,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <Avatar sx={{ m: 1, bgcolor: blue[500] }}>
        {store.userStore.User.name[0]}{store.userStore.User.name[0]}
        </Avatar>
        <Typography component="h1" variant="h5">
          Минибанк
        </Typography>


      </Box>
    </Container>
  )
}

const MainPageObserver = observer(MainPage)

export default MainPageObserver;