import { GetComputers, DeploySoftware, DeployOffice, DeployPrinters, HostName, DeploySophos, DeploySentinel } from "../wailsjs/go/backend/App"
import { useEffect, useState } from "react";
import { backend } from "../wailsjs/go/models"
import {Menu,Item,useContextMenu,Submenu } from "react-contexify";
import "react-contexify/dist/ReactContexify.css";

const MENU_ID = "computer-name";

function App() {
    
    type Computer = {
        name: string;
        ip: string;
        status: string;
      };
    
      const { show } = useContextMenu({ id: MENU_ID });
      const [selectedComputer, setSelectedComputer] = useState<backend.Computer | null>(null);
      // const [targetPC, setTargetPC] = useState<backend.Computer[]>([]);
      const [computers, setComputers] = useState<backend.Computer[]>([]);
      const [hostName, setHostName] = useState<string>("")

      useEffect(() => {
        HostName()
        .then((name: string) => setHostName(name)) // ✅ explicitly show it returns a string
        .catch(err => console.error("HostName error:", err));
      }, []);

    useEffect(() => {
        GetComputers().then(setComputers).catch(console.error);
      }, []);

    const handleContextMenu = (e: React.MouseEvent, comp: backend.Computer) => {
        e.preventDefault,
        setSelectedComputer(comp),

        show({ event: e.nativeEvent })
        
    }

    const handleCommand = ( label: string ) => {
        alert(`you have selected "${label}" for ${selectedComputer?.Name}`);
        
    }

    const handleRefresh = () => {
        GetComputers().then(setComputers).catch(console.error)
    }

    // func for selecting pc and passing it to the deploy software func
    const handleOffice = () => {
      if (selectedComputer?.Name) {
        DeployOffice(selectedComputer.Name)
      } 
      
    }

    const handleSentinel = () => {
        if (selectedComputer?.Name) {
            DeploySentinel(selectedComputer.Name)
        }

    }

    const handlePrinters = () => {
      if (selectedComputer?.Name) {
        DeployPrinters(selectedComputer.Name)
        .then(() => alert("Deployment started"))
      .catch(err => alert("Error: " + err));
      } 
      
    }

    const handleSophos = () => {
        if (selectedComputer?.Name) {
            DeploySophos(selectedComputer.Name)
        }

    }

  return (
    <div>
      <h1>Computer List <button onClick={handleRefresh} className="refresh-button" style={{color:"green", cursor: "pointer" }}>REFRESH</button></h1><p>Running on {hostName}</p>
      { computers.length === 0 ? (
       <p>Loading, please wait.</p>
      ) : (
      <ul>
        {computers.map((c, i) => (
          <li 
          key={i}
          onContextMenu={(e) => handleContextMenu(e, c)}
          style={{ cursor: "pointer", padding: "5px", textDecoration: "underline"}}
          className="computer-row"
          >
            {c.Name} - {c.IP} - {c.Status}</li>
        ))}
        
      </ul>
      )}
        <Menu id={MENU_ID}>
            
            <Submenu label="Software deployment">
            <Item onClick={() => handleOffice()}>Office 365</Item>
            <Item onClick={() => handlePrinters()}>Printer install</Item>
                <Item onClick={() => handleSophos()}>Sophos</Item>
                <Item onClick={() => handleSentinel()}>Sentinel One</Item>
            </Submenu>
            <Item onClick={() => handleCommand("Command 2")}>place holder bullshit</Item>
            <Item onClick={() => handleCommand("Command 3")}>other placeholder</Item>
        </Menu>

    </div>
  );
}

export default App;
