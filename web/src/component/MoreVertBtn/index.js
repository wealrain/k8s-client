import IconButton from '@mui/material/IconButton';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import * as React from 'react';
import { styled, alpha } from '@mui/material/styles';
import { useContext } from 'react';
import { HttpContext } from '../../http';

const StyledMenu = styled((props) => (
    <Menu
      elevation={0}
      anchorOrigin={{
        vertical: 'bottom',
        horizontal: 'right',
      }}
      transformOrigin={{
        vertical: 'top',
        horizontal: 'right',
      }}
      {...props}
    />
  ))(({ theme }) => ({
    '& .MuiPaper-root': {
      borderRadius: 6,
      marginTop: theme.spacing(1),
      minWidth: 180,
      color:
        theme.palette.mode === 'light' ? 'rgb(55, 65, 81)' : theme.palette.grey[300],
      boxShadow:
        'rgb(255, 255, 255) 0px 0px 0px 0px, rgba(0, 0, 0, 0.05) 0px 0px 0px 1px, rgba(0, 0, 0, 0.1) 0px 10px 15px -3px, rgba(0, 0, 0, 0.05) 0px 4px 6px -2px',
      '& .MuiMenu-list': {
        padding: '4px 0',
      },
      '& .MuiMenuItem-root': {
        '& .MuiSvgIcon-root': {
          fontSize: 18,
          color: theme.palette.text.secondary,
          marginRight: theme.spacing(1.5),
        },
        '&:active': {
          backgroundColor: alpha(
            theme.palette.primary.main,
            theme.palette.action.selectedOpacity,
          ),
        },
      },
    },
  }));

const ITEM_HEIGHT = 48;

function MoreVertBtn({items,color}) {
    const [anchorEl, setAnchorEl] = React.useState(null);
    const httpContext = useContext(HttpContext);
    const handleClick = (event) => {
        setAnchorEl(event.currentTarget);
    };
    
    const handleClose = () => {
        setAnchorEl(null);
    };

    const afterHandler = (msg,type)  => {
        httpContext.showMsg(msg,type) ;
    }

    
    return (
        <div>
        <IconButton
            onClick={handleClick}
        >
            <MoreVertIcon sx={{color:(color?color:"gray")}}/>
        </IconButton>
        <StyledMenu
            id="long-menu"
            anchorEl={anchorEl}
            keepMounted
            open={Boolean(anchorEl)}
            onClose={handleClose}
            PaperProps={{
            style: {
                maxHeight: ITEM_HEIGHT * 4.5,
                width: '20ch',
            },
            }}
        >
            {items && items.map((item, index) => (
                <MenuItem key={item.title} onClick={()=>{
                    if(item.handle){
                        let promise = item.handle();
                        if(promise && promise.then){
                            promise.then((res)=>{
                                afterHandler(res.msg,res.type)
                            },(err)=>{
                                afterHandler(err.msg,err.type)
                            })
                        }
                         
                    }
                    handleClose()
                }}>
                    {item.icon}
                    {item.title}
                </MenuItem>
            ))}
            
        </StyledMenu>
        </div>
    );
}

export default MoreVertBtn;