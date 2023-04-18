import { 
    DeleteOutlineOutlined,
    EditOutlined,
    FitScreenOutlined,
    TerminalOutlined,
    ArticleOutlined,
    RestartAltOutlined,
    OpenInNewOutlined,
    PauseCircleOutlineOutlined,
    NotStartedOutlined

} from '@mui/icons-material';
import opHttp from '../http/op';

function createDelete(kind,namespace,name) {
    return {
        title: 'Delete',
        icon: <DeleteOutlineOutlined />,
        handle: () => {
            console.log(kind,namespace,name);
        }
    }
}

function createEdit(kind,namespace,name) {
    return {
        title: 'Edit',
        icon: <EditOutlined />,
        handle: () => {
            window.open(`/edit/${kind}/${namespace}/${name}`,"_blank");
        }
    }
}

function createScale() {
    return {
        title: 'Scale',
        icon: <FitScreenOutlined />,
        handle: () => {
            console.log('scale');
        }
    }
}

function createRestart() {
    return {
        title: 'Restart',
        icon: <RestartAltOutlined />,
        handle: () => {
            console.log('restart');
        }
    }
}

function createLog(kind,namespace,name) {
    return {
        title: 'Log',
        icon: <ArticleOutlined />,
        handle: () => {
            window.open(`/log/${kind}/${namespace}/${name}`,"_blank");
        }
    }
}

function createShell() {
    return {
        title: 'Shell',
        icon: <TerminalOutlined />,
        handle: () => {
            console.log('shell');
        }
    }
}

function createAttach() {
    return {
        title: 'Attach',
        icon: <OpenInNewOutlined />,
        handle: () => {
            console.log('attach');
        }
    }
}

function createPause(namespace, name) {
    return {
        title: 'Pause',
        icon: <PauseCircleOutlineOutlined />,
        handle: () => {
            return opHttp.pauseDeployment(namespace, name);
        }
    }
}

function createResume(namespace, name) {
    return {
        title: 'Resume',
        icon: <NotStartedOutlined />,
        handle: () => {
           return opHttp.resumeDeployment(namespace, name);
        }
    }
}



export default {
    createDelete,
    createEdit,
    createScale,
    createRestart,
    createLog,
    createShell,
    createAttach,
    createPause,
    createResume
}
