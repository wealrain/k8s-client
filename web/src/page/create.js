import { 
    Box, 
    AppBar, 
    Toolbar,
    Typography,
    Snackbar,
    Alert,
    TextField,
    OutlinedInput,
    InputLabel,
    MenuItem,
    FormControl,
    Select,
    Button,
} from '@mui/material';
import YamlEditor from '../component/YamlEditor';
import { useEffect,useState,forwardRef } from 'react';
import listHttp from '../http/list';
import op from '../http/op';
import parser from "js-yaml";

const ITEM_HEIGHT = 48;
const ITEM_PADDING_TOP = 8;
const MenuProps = {
  PaperProps: {
    style: {
      maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
      width: 250,
    },
  },
};

const kinds = [
    'pod',
    'deployment',
    'service',
    'ingress',
    'configmap',
    'statefulset',
];

function NameInput(props){
    const {name,setName} = props;
    const handleChange = (event) => {
        const {
            target: { value },
        } = event;
        setName(value);
    };
    return (
            <TextField 
            id="outlined-basic" 
            label="Name" 
            variant="outlined" 
            size='small'
            onChange={handleChange}
            sx={{ mr:1}}
            />
    )
}

function NamespaceSelect(props) {
    const {namespace,setNamespace,namespaces} = props;
  
    const handleChange = (event) => {
      const {
        target: { value },
      } = event;
      setNamespace(value);
    };

    return (
        <div>
        <FormControl sx={{ mr:1, width: 200 }} size='small'>
          <InputLabel id="namespace-label">Namespace</InputLabel>
          <Select
            labelId="namespace-label"
            id="namespace"
            value={namespace}
            onChange={handleChange}
            input={<OutlinedInput label="Namespace" />}
            MenuProps={MenuProps}
          >
            {namespaces.map((namespace) => (
              <MenuItem
                key={namespace}
                value={namespace}
              >
                {namespace}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      </div>
    )
}

function KindSelect(props){
    const {kind,setKind} = props;
  
    const handleChange = (event) => {
      const {
        target: { value },
      } = event;
      setKind(value);
    };

    return (
        <div>
        <FormControl sx={{ mr:1, width: 200 }} size='small'>
          <InputLabel id="kind-label">Kind</InputLabel>
          <Select
            labelId="kind-label"
            id="kind"
            value={kind}
            onChange={handleChange}
            input={<OutlinedInput label="Kind" />}
            MenuProps={MenuProps}
          >
            {kinds.map((kind) => (
              <MenuItem
                key={kind}
                value={kind}
              >
                {kind}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      </div>
    )
}


const MuiAlert = forwardRef(function MuiAlert(props, ref) {
  return <Alert elevation={6} ref={ref} variant="filled" {...props} />;
});

export default function Create() {
    
    const [currentYaml, setCurrentYaml] = useState('');
    const [type ,setType] = useState('success');
    const [msg,setMsg] = useState('');
    const [openAlert,setOpenAlert] = useState(false);
    const [kind,setKind] = useState('pod');
    const [namespace,setNamespace] = useState('');
    const [namespaces,setNamespaces] = useState([]);
    const [name,setName] = useState('');

    const handleCloseAlert = () => setOpenAlert(false);
    const showMsg = (msg,type) => {
        setOpenAlert(true);
        setMsg(msg);
        setType(type);
    }

    function sendYaml(){
        // 校验yaml
        try {
            let value = parser.load(currentYaml);
            // 发送请求
            op.putResource(kind,namespace,name,value).then(res => {
                showMsg("创建成功",'success');
            }).catch(err => {
                if(!err.msg) err.msg = 'unknow error';
                showMsg(err.msg,'error');
            })
        } catch (error) {
            showMsg("yaml 错误",'error');
            return;
        }
    }

    function handleYamlChange(value) {
        setCurrentYaml(value);
    }

    const fetchNamespaces = async () => {
        const res = await listHttp.listNamespace();
        setNamespaces(res);
        setNamespace(res[0]);
    }

    useEffect(() => {
         fetchNamespaces();
    }, [])

    return (
        <Box sx={{ display: 'flex' }}>
      <AppBar
        position="fixed"
        sx={{
            backgroundColor: '#fff',
            color: '#000',
            zIndex: (theme) => theme.zIndex.drawer + 1,
        }}
      >
        <Toolbar>
        <Typography variant="h6" noWrap component="div" >
           资源编辑器： 
        </Typography>
        <KindSelect kind={kind} setKind={setKind}/>
        <NamespaceSelect namespace={namespace} setNamespace={setNamespace} namespaces={namespaces}/>
        <NameInput name={name} setName={setName}/>
        <Button variant="contained" size='large' onClick={sendYaml}>Create</Button>
        </Toolbar>
      </AppBar>
       
      <Box
        component="main"
        sx={{ flexGrow: 1, }}
      >
        <Toolbar />
        <YamlEditor onChange={handleYamlChange}/>
       
      </Box>
      <Snackbar open={openAlert} autoHideDuration={3000} onClose={handleCloseAlert}>
            <MuiAlert onClose={handleCloseAlert} severity={type} sx={{ width: '100%' }}>
                    {msg}
            </MuiAlert>
            </Snackbar>
    </Box>
    )
}