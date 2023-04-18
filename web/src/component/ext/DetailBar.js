import { 
    Box, 
    AppBar,
    Toolbar,
    Typography,
    IconButton,
} from "@mui/material";
import MoreVertBtn from '../MoreVertBtn';
export default function DetailBar(props) {
    return (
        <Box sx={{width:"100%"}}>
              <AppBar sx={{
                position:"absolute",
              }}>
      <Toolbar >
        <Typography variant="h6" noWrap component="div" sx={{ flexGrow: 1 }}>
             {props.title}   
        </Typography>
        <MoreVertBtn color={"#fff"} items={props.moreHandles}/>
       </Toolbar>
    </AppBar>
        </Box>
    )
  }