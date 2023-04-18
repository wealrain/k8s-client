import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import clusteHttp from '../../http/cluster'
import { HttpContext } from '../../http';

export default function AddCluster(props) {
  const { open, setOpen } = props;
  const [name, setName] = React.useState('');
  const [config, setConfig] = React.useState('');

  const httpContext = React.useContext(HttpContext);
 
  const handleAddCluster = () => {
        clusteHttp.createCluster({
            name,
            config
        }).then((res)=>{
          httpContext.showMsg(res.msg,res.type) ;
          setOpen(false);
        }).catch((err)=>{
          console.log(err);
          httpContext.showMsg(err.msg,err.type) ;
        })
  }
  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Add Cluster</DialogTitle>
        <DialogContent>
          <DialogContentText>
            please add cluster width name and config content
          </DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            id="name"
            label="Name"
            type="text"
            fullWidth
            variant="standard"
            onChange={(e)=>{
                setName(e.target.value)
            }}
          />
          <TextField
            id="standard-multiline-flexible"
            label="Config"
            multiline
            fullWidth
            variant="standard"
            onChange={(e)=>{
                setConfig(e.target.value)
            }}

        />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleAddCluster}>Submit</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}