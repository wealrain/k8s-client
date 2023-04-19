import logo from './logo.svg';
import './App.css';
import router from './router';
import { RouterProvider } from 'react-router-dom';
import { HttpProvider } from './http';
import { currentCluster } from './store/cluster';
import ChooseCluster from './component/ChooseCluster';
import * as React from 'react';

export const AppContext = React.createContext(null);

function App() {
  const [open, setOpen] = React.useState(false);
  const [cluster, setCluster] = React.useState(null);
  const [namespace,setNamespace] = React.useState(null);

  React.useEffect(() => {
    setCluster(currentCluster());
    if(!cluster){
      setOpen(true)
    } else {
      setOpen(false)
    }
  }, [cluster])
  return (
    <HttpProvider>
        <AppContext.Provider value={{cluster,setCluster,namespace,setNamespace}}>
        <RouterProvider router={router} />
        <ChooseCluster open={open} setOpen={setOpen}/>
        </AppContext.Provider>
    </HttpProvider>
  );
}

export default App;
