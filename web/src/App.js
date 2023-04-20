import './App.css';
import router from './router';
import { RouterProvider } from 'react-router-dom';
import { HttpProvider } from './http';
import { currentCluster } from './store/cluster';
import { setUserToken, userToken } from './store/token';
import ChooseCluster from './component/ChooseCluster';
import * as React from 'react';
import listHttp from './http/list';
import Login from './page/login';

export const AppContext = React.createContext(null);

function App() {

  const [open, setOpen] = React.useState(false);
  const [cluster, setCluster] = React.useState(null);
  const [namespace,setNamespace] = React.useState(null);
  const [namespaces,setNamespaces] = React.useState([]);

  const [token,setToken] = React.useState(null);

  const fetchNamespace = async () => {
    let result = await listHttp.listNamespace();

    console.log(result);
    setNamespace(result[0]);
    setNamespaces(result);
  } 
  
  React.useEffect(() => {
    setToken(userToken());
    setCluster(currentCluster());
    if(!cluster){
      setOpen(true)
    } else {
      fetchNamespace();
      setOpen(false)
    }
  }, [cluster,token])
  return (
     
    <HttpProvider>
        <AppContext.Provider value={{cluster,setCluster,namespace,setNamespace,namespaces,setNamespaces,setToken}}>
        {
          token ?
            <>
            <ChooseCluster open={open} setOpen={setOpen}/>
              {
                cluster && namespace  && <RouterProvider router={router} />
              }
            </>
          : <Login />
        }
        
        
        </AppContext.Provider>
    </HttpProvider>
     
  );
}

export default App;
