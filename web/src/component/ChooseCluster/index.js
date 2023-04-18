import * as React from 'react';
import {
    AppBar,
    Button,
    Dialog,
    MenuItem,
    Slide,
    Toolbar,
    Typography,
    Select,
    Box,
    IconButton,
} from '@mui/material';
import clusterHttp from '../../http/cluster'
import AddCluster from '../Setting/AddCluster';
import { setCurrentCluster } from '../../store/cluster';

const Transition = React.forwardRef(function Transition(props, ref) {
  return <Slide direction="up" ref={ref} {...props} />;
});

export default function ChooseCluster(props) {
  const { open, setOpen } = props;
  const [openAddCluster, setOpenAddCluster] = React.useState(false);
  const [cluster, setCluster] = React.useState('');
  const [clusterList, setClusterList] = React.useState([]);

  const handleChange = (event) => {
    console.log(event.target.value);
    setCluster(event.target.value);
    setCurrentCluster(event.target.value);

    setOpen(false);
  };

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const getClusterList = async () => {
    const res = await clusterHttp.listClusters();
    let items = res.map((item)=>{
      return {
        name:item.name,
        id: item.id
      }
    })
    setClusterList(items);
  }

  React.useEffect(() => {
    getClusterList();
  }, [open]);

  return (
    <div>
      <Dialog
        fullScreen
        open={open}
        onClose={handleClose}
        TransitionComponent={Transition}
      >
        <AppBar sx={{ position: 'relative' }}>
          <Toolbar>
            <Typography sx={{ ml: 2, flex: 1 }} variant="h6" component="div">
              please choose a cluster
            </Typography>
            <Button autoFocus color="inherit" onClick={()=>{
                setOpenAddCluster(true)
            }}>
              add
            </Button>

            <Button autoFocus color="inherit" onClick={handleClose}>
              save
            </Button>
          </Toolbar>
        </AppBar>
        {/* <Toolbar /> */}
        <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center',flexDirection:"column" }}>
        <Typography sx={{ mt: 4, ml: 4,marginBottom:4}} variant="h6" component="div">
        Please select a cluster below as your default cluster. If there is no selection below, please use the add button to add a cluster as your default cluster.
        </Typography>
        <Select
          sx={{ width: 300, marginBottom: 4 }}
          labelId="demo-simple-select-filled-label"
          id="demo-simple-select-filled"
          value={cluster}
          onChange={handleChange}
        >
          <MenuItem value="">
            <em>None</em>
          </MenuItem>
          {
            clusterList.map((item)=>{
              return <MenuItem value={item.id} key={item.id}>{item.name}</MenuItem>
            })
          }
        </Select>
        </Box>
        
      </Dialog>
      <AddCluster open={openAddCluster} setOpen={setOpenAddCluster} />
    </div>
  );
}