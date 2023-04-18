import {ReplaySubject, Subject, debounce, takeUntil} from 'rxjs'
import {Terminal} from 'xterm'
import {FitAddon} from 'xterm-addon-fit'
import {useEffect, useRef, useState} from 'react'
import {useParams} from 'react-router-dom'
import {execShell} from '../../http/shell'

function Shell() {

    let namespace = null;
    let pod = null;
    let container = null;

    let term = null;
    let connSubject = new ReplaySubject(100); // 用于接收消息，缓存100条
    let unsubscribe = new Subject();
    let keyEvent = new ReplaySubject(2); // 用于接收按键事件，缓存2条

    let conn = null; // websocket连接
    let connecting = false; // 是否正在连接
    let connected = false; // 是否已连接
    let anchorRef = useRef(null);

    let debounceFit = null;

    function init() {
        // 从url中获取参数
       

        
    }

    async function setupConnection() {
        let {namespace, pod, container} = useParams();
        const {id} = await execShell({
            namespace,
            pod,
            container,
        })

        conn = new SockJS(`/api/sock?${id}`);
        conn.onopen = () => {
            onConnectionOpen(id);
        };
        conn.onmessage = () => {
            onConnectionMessage(id);
        };
        conn.onclose = ()=>{ 
            onConnectionClose(id);
        };
    }

    function onConnectionOpen(id) {
        const startData = {
            Op:'bind',
            SessionId: id
        }
        conn.send(JSON.stringify(startData))
        connSubject.next(startData)

        connected = true
        connecting = false
        
        onTerminalResize()
        term.focus()


    }
    
    function initXterm() {
        if(connSubject) {
            connSubject.complete();
            connSubject = new ReplaySubject(100);
        }

        if(term) {
            term.dispose();
        }

        term = new Terminal({
            fontSize: 14,
            fontFamily: 'Consolas, "Courier New", monospace',
            bellStyle: 'sound',
            cursorBlink: true,
        });

        const fitAddon = new FitAddon();
        term.loadAddon(fitAddon);
        term.open(anchorRef.current);

        debounceFit = debounce(() => {
            fitAddon.fit();
            // todo 更新状态，触发重绘
        }, 100);
        debounceFit();

        window.addEventListener('resize', ()=> debounceFit());

        // 获取消息直到取消订阅
        connSubject.pipe(takeUntil(unsubscribe)).subscribe((frame) => {
            handleConnectionMessage(frame);
        });

        term.onData((data) => {
            onTermianlData(data);   
        });
        
        term.onResize(() => onTerminalResize());
        term.onKey((e) => keyEvent.next(e.domEvent));
    }


    function handleConnectionMessage(frame) {
    }

    function onTermianlData(data) {
        if(connected) {

        }
    }

    function onTerminalResize() {}
}