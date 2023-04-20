import React, { useState,forwardRef,useContext } from "react";
import {
  Avatar,
  Button,
  CssBaseline,
  TextField,
  Link,
  Grid,
  Box,
  Typography,
  Container,
  createTheme,
  ThemeProvider,
  Snackbar,
  Alert,
} from "@mui/material";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import tokenHttp from "../http/token";
import { setUserToken } from "../store/token";
import { AppContext } from "../App";

const theme = createTheme();

const MuiAlert = forwardRef(function MuiAlert(props, ref) {
    return <Alert elevation={6} ref={ref} variant="filled" {...props} />;
  });

export default function SignIn() {
  const {setToken} = useContext(AppContext)
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [type ,setType] = useState('success');
  const [msg,setMsg] = useState('');
  const [openAlert,setOpenAlert] = useState(false);
  const handleCloseAlert = () => setOpenAlert(false);
  const showMsg = (msg,type) => {
        setOpenAlert(true);
        setMsg(msg);
        setType(type);
    }
  const handleSubmit = (e) => {
    e.preventDefault();
    // handle login logic here
    tokenHttp.createToken(username, password).then((res) => {
        showMsg("login success",'success');
        // 保存token
        setUserToken(res.token);
        setToken(res.token);
    }).catch((err) => {
        if(!err.msg) err.msg = 'unknow error';
        showMsg(err.msg,'error');
    });
  };

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <Box component="form" onSubmit={handleSubmit} sx={{ mt: 1 }}>
            <TextField
              margin="normal"
              required
              fullWidth
              id="username"
              label="Username"
              name="username"
              autoComplete="username"
              autoFocus
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Sign In
            </Button>
            <Grid container>
              <Grid item xs>
                <Link href="#" variant="body2">
                  Forgot password?
                </Link>
              </Grid>
              <Grid item>
                <Link href="#" variant="body2">
                  {"Don't have an account? Sign Up"}
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
         
      </Container>
      <Snackbar open={openAlert} autoHideDuration={3000} onClose={handleCloseAlert}>
            <MuiAlert onClose={handleCloseAlert} severity={type} sx={{ width: '100%' }}>
                    {msg}
            </MuiAlert>
            </Snackbar>
    </ThemeProvider>
  );
}
