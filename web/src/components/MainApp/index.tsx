import React, {FC} from "react";
import ContactList from "../ContactList";

const MainApp: FC = function (props) {
    return (
        <div className="main">
            <ContactList/>
        </div>
    );
}

export default MainApp;