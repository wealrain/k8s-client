import {createBrowserRouter} from 'react-router-dom';
import Layout from '../component/Layout';
import Home from '../page/home';
import PodList from '../page/podList';
import DeploymentList from '../page/deploymentList';
import StatefulList from '../page/statefulSetList';
import ReplicaSetList from '../page/replicaSetList';
import ConfigMapList from '../page/configmapList';
import SecretList from '../page/secretList';
import ServiceList from '../page/serviceList';
import IngressList from '../page/ingressList';
import EndpointList from '../page/endpointList';
import EventList from '../page/eventList';
import Edit from '../page/edit';
import Log from '../page/log';
import Create from '../page/create';

let router = createBrowserRouter([
    {
        path: '/create',
        Component: Create,
    },
    {
        path: '/edit/:kind/:namespace/:name',
        Component: Edit,
    },
    {
        path: '/log/:kind/:namespace/:name',
        Component: Log,
    },
    {
        path: '/',
        Component: Layout,
        children: [
            {
                path: '/',
                Component: Home
            },
            {
                path: '/workload/pods',
                Component: PodList
            },
            {
                path: '/workload/deployments',
                Component: DeploymentList
            },
            {
                path: '/workload/statefulsets',
                Component: StatefulList
            },
            {
                path: '/workload/replicasets',
                Component: ReplicaSetList
            },
            {
                path: '/config/configmaps',
                Component: ConfigMapList
            },
            {
                path: '/config/secrets',
                Component: SecretList
            },
            {
                path: '/network/services',
                Component: ServiceList
            },
            {
                path: '/network/ingresses',
                Component: IngressList
            },
            {
                path: '/network/endpoints',
                Component: EndpointList
            },
            {
                path: '/events',
                Component: EventList
            }
        ]
    }
]);


export default router;