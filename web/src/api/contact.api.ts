import {Contact} from "../types/contact";
import {getClient} from "../services/http.service";
import {AxiosResponse} from "axios";

export function getContactsList(): Promise<Contact[]> {
    return new Promise((resolve, reject) => {
        getClient()
            .get('api/contacts')
            .then(({data}: AxiosResponse<{ id: string, name: string, age: number, phone: string }[]>) => {
                resolve(data);
            }).catch(reject)
    });
}

export function storeContact(contact: Contact): Promise<void> {
    const form = new FormData();
    for(const key in contact) {
        // @ts-ignore
        form.append(key, contact[key]);
    }
    return getClient().post('api/contacts', form);
}

export function deleteContact(contact: Contact): Promise<void> {
    return getClient().delete(`api/contacts/${contact.id}`);
}
