import React, {ChangeEvent, FC, useCallback, useEffect, useState} from "react";
import {deleteContact, getContactsList, storeContact} from "../../api/contact.api";
import {Contact} from "../../types/contact";

const ContactList: FC = function () {
    const [contacts, setContacts] = useState<Contact[]>([]);
    const [nContact, setNContact] = useState<{ name: string, age: number, phone: string }>({
        name: '',
        age: 0,
        phone: ''
    });
    const getContacts = useCallback(async () => {
        setContacts(await getContactsList());
    }, [setContacts]);
    const delContact = useCallback((contact: Contact) => {
        return async () => {
            if (confirm('Are you sure?')) {
                await deleteContact(contact).then(() => getContacts());
            }
        }
    }, []);

    useEffect(() => {
        getContacts().then(r => r);
    }, []);


    const updateField = (e: ChangeEvent<HTMLInputElement>): void => {
        setNContact({...nContact, [e.target.name]: e.target.value})
    };

    const saveContact = async () => {
        await storeContact({id: '', ...nContact})
            .then(() => {
                setNContact({name: '', age: 0, phone: ''});
                getContacts()
            });
    };

    return (
        <div className="main">
            <table>
                <thead>
                <tr>
                    <th>Name</th>
                    <th>Age</th>
                    <th>Phone</th>
                    <th>Action</th>
                </tr>
                </thead>
                <tbody>
                {contacts?.map(contact => {
                    return (
                        <tr key={contact.id}>
                            <td>{contact.name}</td>
                            <td>{contact.age}</td>
                            <td>{contact.phone}</td>
                            <td>
                                <button onClick={delContact(contact)}>DEL</button>
                            </td>
                        </tr>
                    )
                })}
                <tr>
                    <td><input type="text" name="name" value={nContact.name} onChange={updateField}/></td>
                    <td><input type="number" name="age" value={nContact.age} onChange={updateField}/></td>
                    <td><input type="text" name="phone" value={nContact.phone} onChange={updateField}/></td>
                    <td><button type="button" onClick={saveContact}>Save</button></td>
                </tr>
                </tbody>
            </table>
        </div>
    );
}

export default ContactList;