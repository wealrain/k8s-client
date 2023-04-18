import logo from './logo.svg';
import './App.css';
import router from './router';
import { RouterProvider } from 'react-router-dom';
import { HttpProvider } from './http';
import { currentCluster } from './store/cluster';
import ChooseCluster from './component/ChooseCluster';
import * as React from 'react';


function App() {
  const [open, setOpen] = React.useState(false);
  let cluster = currentCluster()
  React.useEffect(() => {
    if(!cluster){
      setOpen(true)
    } else {
      setOpen(false)
    }
  }, [cluster])
  return (
    <HttpProvider>
        <RouterProvider router={router} />
        <ChooseCluster open={open} setOpen={setOpen}/>
    </HttpProvider>
  );
}

export default App;
