import './App.css';
import router from './router';
import { RouterProvider } from 'react-router-dom';
import { HttpProvider } from './http';
import { currentCluster } from './store/cluster';
import ChooseCluster from './component/ChooseCluster';
import * as React from 'react';
import listHttp from './http/list';

export const AppContext = React.createContext(null);

function App() {
  const [open, setOpen] = React.useState(false);
  const [cluster, setCluster] = React.useState(null);
  const [namespace,setNamespace] = React.useState(null);
  const [namespaces,setNamespaces] = React.useState([]);

  const fetchNamespace = async () => {
    let result = await listHttp.listNamespace();
    setNamespace(result[0]);
    setNamespaces(result);
  } 
  
  React.useEffect(() => {
    setCluster(currentCluster());
    if(!cluster){
      setOpen(true)
    } else {
      fetchNamespace();
      setOpen(false)
    }
  }, [cluster])
  return (
    <HttpProvider>
        <AppContext.Provider value={{cluster,setCluster,namespace,setNamespace,namespaces,setNamespaces}}>
        <ChooseCluster open={open} setOpen={setOpen}/>
        {
          cluster && namespace  && <RouterProvider router={router} />
        }
        </AppContext.Provider>
    </HttpProvider>
  );
}

export default App;
