import { 
    Box, 
    AppBar, 
    Toolbar,
    Typography,
    SpeedDial,
    SpeedDialIcon,
    SpeedDialAction,
    Snackbar,
    Alert,
} from '@mui/material';
import {
  FileCopyOutlined, 
  SaveAsOutlined,
  LocalPrintshopOutlined,
  ShareOutlined, 
} from '@mui/icons-material';
import { useParams } from 'react-router-dom';
import { githubLight } from "@uiw/codemirror-theme-github";
import { useEffect,useState,forwardRef } from 'react';
import log from '../http/log';
import CodeMirror from "@uiw/react-codemirror";
import SockJS from 'sockjs-client';

 

const actions = [
  { icon: <FileCopyOutlined />, name: 'download',handleCreate: (kind,namespace,name,value,showMsg) =>(
    function(){
         
        log.getLog(namespace,name,"").then((res)=>{
            const blob = new Blob([res.data]);
            const fileName = res.headers["content-disposition"].split(";")[1].split("filename=")[1];
            if ('download' in document.createElement("a")) {
                const link = document.createElement("a");
                link.download = fileName;
                link.style.display = 'none';
                link.href = URL.createObjectURL(blob);
                document.body.appendChild(link);
                link.click();
                URL.revokeObjectURL(link.href);
                document.body.removeChild(link);
            } else {
                navigator.msSaveBlob(blob, fileName);
            }
        },(err)=>{
            showMsg(err.response.data.error,'error');
            }
        );

    }
  )},
  { icon: <SaveAsOutlined />, name: 'Save',handleCreate: (kind,namespace,name,value,showMsg) =>(
    function(){
       

    }
  ) },
  { icon: <LocalPrintshopOutlined />, name: 'Print',handleCreate: (kind,namespace,name,value) =>(
    function(){
      const w = window.open('about:blank');
      w.document.write(value);
      w.document.close();
    })
  },
  { icon: <ShareOutlined />, name: 'Share',handleCreate: (kind,namespace,name,value) =>(
    function(){
      const url = window.location.href;
      navigator.clipboard.writeText(url).then(function() {
        alert('已复制到剪贴板');
      }, function(err) {
        console.error('Could not copy text: ', err);
      });
    })      
  },
];

const MuiAlert = forwardRef(function MuiAlert(props, ref) {
  return <Alert elevation={6} ref={ref} variant="filled" {...props} />;
});
let sock;

export default function Edit() {
    
    const parmas = useParams();
    const { kind, namespace, name,container } = parmas;
    const [value, setValue] = useState("");
    const [open, setOpen] = useState(false);
    const [type ,setType] = useState('success');
    const [msg,setMsg] = useState('');
    const [openAlert,setOpenAlert] = useState(false);

    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);
    const handleCloseAlert = () => setOpenAlert(false);
    const showMsg = (msg,type) => {
        setOpenAlert(true);
        setMsg(msg);
        setType(type);
    }


    

    async function getWSConection(){
        let session = await log.getLogSession(namespace,name,container?container:"")
        let sock = new SockJS(`http://localhost:8080/logs`);
        sock.onopen = function() {
            const msg = {
                op:"bind",
                sessionId:session.sessionId,
            }
            
            sock.send(JSON.stringify(msg));
        } 
        sock.onmessage = function(e) {
            setValue((preValue) => preValue.concat(e.data));
        }
        sock.onclose = function() {
            console.log("close");
        }
        return sock;
         
    }

    useEffect(() => {
        setValue("");
        
        if(!sock) {
            sock = getWSConection();
          console.log("sock",sock);
          }
        
        // return () => {
        //     sock.then((sock)=>{
        //         sock.close();
        //     })
        // }
       
    }, [namespace,name,container])

    return (
        <Box sx={{ display: 'flex' }}>
      <AppBar
        position="fixed"
      >
        <Toolbar>
        <Typography variant="h6" noWrap component="div" >
           日志编辑器：{parmas.kind}/{parmas.namespace}/{parmas.name}  
        </Typography>
        </Toolbar>
      </AppBar>
       
      <Box
        component="main"
        sx={{ flexGrow: 1, }}
      >
        <Toolbar />
        <CodeMirror
            width="100%"
            value={value}
            theme={githubLight}
        />
        <SpeedDial
        ariaLabel="SpeedDial controlled open example"
        sx={{ position: 'fixed', bottom: 40, right: 16 }}
        icon={<SpeedDialIcon />}
        onClose={handleClose}
        onOpen={handleOpen}
        open={open}
      >
        {actions.map((action) => (
          <SpeedDialAction
            key={action.name}
            icon={action.icon}
            tooltipTitle={action.name}
            onClick={action.handleCreate(kind,namespace,name,value,showMsg)}
          />
        ))}
      </SpeedDial>
      </Box>
      <Snackbar open={openAlert} autoHideDuration={3000} onClose={handleCloseAlert}>
            <MuiAlert onClose={handleCloseAlert} severity={type} sx={{ width: '100%' }}>
                    {msg}
            </MuiAlert>
            </Snackbar>
    </Box>
    )
}