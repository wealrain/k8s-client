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
import YamlEditor from '../component/YamlEditor';
import { useEffect,useState,forwardRef } from 'react';
import op from '../http/op';
import parser from "js-yaml";

const actions = [
  { icon: <FileCopyOutlined />, name: 'Copy',handleCreate: (kind,namespace,name,value) =>(
    function(){
      navigator.clipboard.writeText(value).then(function() {
        alert('已复制到剪贴板');
      }, function(err) {
        console.error('Could not copy text: ', err);
      });
    }
  )},
  { icon: <SaveAsOutlined />, name: 'Save',handleCreate: (kind,namespace,name,value,showMsg) =>(
    function(){
       value = parser.load(value);
       op.putResource(kind,namespace,name,value).then((res)=>{
          showMsg('保存成功','success');
        },(err)=>{
          showMsg(err.response.data.error,'error');
        });

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

export default function Edit() {
    const parmas = useParams();
    const { kind, namespace, name } = parmas;
    const [yaml, setYaml] = useState('');
    const [currentYaml, setCurrentYaml] = useState('');
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

    async function getYaml() {
        const result = await op.getResource(kind, namespace, name);
        // 将json转换为yaml
        result.metadata.managedFields = undefined;
        const yaml = parser.dump(result);
        setYaml(yaml);
        setCurrentYaml(yaml);
    }
 
    function handleYamlChange(value) {
        setCurrentYaml(value);
    }

    useEffect(() => {
        getYaml();
    }, [kind, namespace, name])

    return (
        <Box sx={{ display: 'flex' }}>
      <AppBar
        position="fixed"
      >
        <Toolbar>
        <Typography variant="h6" noWrap component="div" >
           资源编辑器：{parmas.kind}/{parmas.namespace}/{parmas.name}  
        </Typography>
        </Toolbar>
      </AppBar>
       
      <Box
        component="main"
        sx={{ flexGrow: 1, }}
      >
        <Toolbar />
        <YamlEditor value={yaml} onChange={handleYamlChange}/>
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
            onClick={action.handleCreate(kind,namespace,name,currentYaml,showMsg)}
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