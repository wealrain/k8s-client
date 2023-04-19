import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import Divider from '@mui/material/Divider';
import React,{useEffect} from 'react';
import ExpandLess from '@mui/icons-material/ExpandLess';
import ExpandMore from '@mui/icons-material/ExpandMore';
import Collapse from '@mui/material/Collapse';
import Typography from '@mui/material/Typography';
import { AppBar, Link } from '@mui/material';
import { useNavigate  } from 'react-router-dom';
import { 
  AssuredWorkload,
  Settings, 
  Podcasts,
  EventNote
} from '@mui/icons-material';
import { useLocation } from 'react-router-dom';

const menus = [
  {
    title: 'Workloads',
    icon: <AssuredWorkload />,
    children: [
      {
        title: 'Pods',
        link: '/workload/pods'
      },
      {
        title: 'Deployments',
        link: '/workload/deployments'
      },
      {
        title: 'StatefulSets',
        link: '/workload/statefulsets',
      },
      {
        title: 'ReplicaSets',
        link: '/workload/replicasets',
      },
    ]
  },
  {
    title: 'Config',
    icon: <Settings />,
    children: [
      {
        title: 'ConfigMaps',
        link: '/config/configmaps'
      },
      {
        title: 'Secrets',
        link: '/config/secrets'
      }
    ]
  },
  {
    title: 'Network',
    icon: <Podcasts />,
    children: [
      {
        title: 'Services',
        link: '/network/services'
      },
      {
        title: 'Ingresses',
        link: '/network/ingresses'
      },
      {
        title: 'Endpoints',
        link: '/network/endpoints'
      }
    ]
  },
  {
    title: 'Events',
    icon: <EventNote />,
    link: '/events'
  },
]

function Menu() {
  const [selectedIndex, setSelectedIndex] = React.useState(0);

  // 根据路由设置选中的菜单
  const location = useLocation();
  useEffect(() => {
    const path = location.pathname;
    const index = menus.findIndex(menu => path.toLowerCase().indexOf(menu.title.toLowerCase()) !== -1 );
    if(index !== -1){
      setSelectedIndex(index);
    }
  }, [location.pathname])

  function handleClick(index) {
    setSelectedIndex(index);
  }

  return (
  <div>
    <Divider />
    <List>
      {menus.map((menu, index) => (
        <MenuItem 
          key= {index} 
          item={menu} 
          index={index} 
          selected={selectedIndex === index} 
          handleClick={handleClick} />
      ))}
    </List>
  </div>
  )
}

function MenuItem({item,index,selected,handleClick}) {

  const [open, setOpen] = React.useState(false);

  const navigate = useNavigate();

  useEffect(() => {
    setOpen(selected);
  }, [selected])

  return (
    <>
    <ListItem key={item.title} disablePadding>
    <ListItemButton onClick={()=>{
        handleClick(index)
        if(item.link){
          navigate(item.link)
        }
    }}>
      <ListItemIcon>
        {item.icon}
      </ListItemIcon>
      <ListItemText primary={item.title}/>
      {item.children && (open ? <ExpandLess /> : <ExpandMore />)}
      </ListItemButton>
    </ListItem>
    <Collapse in={open} timeout="auto" unmountOnExit>
        <List component="div" disablePadding>
          {item.children && item.children.map((child, index) => (
            <ListItem key={child.title} sx={{ pl: 4 }}>
              <ListItemButton onClick={()=>{
                  if(child.link){
                    navigate(child.link)
                  }
              }}>
                <ListItemIcon>
                    {child.icon}
                  </ListItemIcon>
                  <ListItemText primary={child.title}/>
                </ListItemButton>
            </ListItem>
          ))}
        </List>
      </Collapse>
    </>
  )
}

 

export default Menu;