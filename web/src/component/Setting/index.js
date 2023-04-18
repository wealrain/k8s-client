import { 
    Drawer,
    Box,
    ListItemButton,
    ListItemText,
    ListItemIcon,
 } from "@mui/material";
import {
    useState,
    createContext,
    useContext,
 } from "react";
import {
    AccountTreeOutlined,
    AutoAwesomeMotionOutlined
 } from "@mui/icons-material";
import General from "./General";
import Clusters from "./Clusters";

function Slide() {
    const {selectedIndex,setSelectedIndex} = useContext(SettingContext);
    
    const handleListItemClick = (event, index) => {
        setSelectedIndex(index);
    };

    const data = [
        {
            id: 1,
            label: 'General',
            icon: <AccountTreeOutlined />,
        },
        {
            id: 2,
            label: 'Clusters',
            icon: <AutoAwesomeMotionOutlined />,
        },
    ]
    return (
        <>
              { 
                data.map((item) => (
                  <ListItemButton
                    key={item.label}
                    sx={{ py: 0, minHeight: 32 }}
                    selected={selectedIndex === item.id}
                    onClick={(event) => handleListItemClick(event, item.id)}
                  >
                    <ListItemIcon sx={{ color: 'inherit' }}>
                      {item.icon}
                    </ListItemIcon>
                    <ListItemText
                      primary={item.label}
                      primaryTypographyProps={{ fontSize: 14, fontWeight: 'medium' }}
                    />
                  </ListItemButton>
                ))
            }
        </>
    )
}

const SettingContext = createContext(null)

function Content() {
    const {selectedIndex} = useContext(SettingContext);
    if(selectedIndex === 1) {
        return <General/>
    } else if(selectedIndex === 2) {
        return <Clusters/>
    }
}

export default function Setting(props) {
    const {open,onClose} = props;
    const [selectedIndex, setSelectedIndex] = useState(1);

    const info = () => (
        <SettingContext.Provider value={{selectedIndex,setSelectedIndex}}>
        <Box
            sx={{ width:1024,height:1,display: 'flex',flexDirection:"row",}}
        >   
            <Box sx={{width:200,height:1,backgroundColor:"rgba(100,100,100,0.1)"}}>
                <Slide/>
            </Box>
            <Box sx={{flexGrow:1,height:200}}>
                <Content/>
            </Box>
        </Box>
        </SettingContext.Provider>
             
             
     ) 

    return (
        <Drawer  
            anchor="left"
            open={open}
            onClose={onClose}
        >
            {info()}
        </Drawer>
    )
}