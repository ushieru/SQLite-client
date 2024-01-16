// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';

export function ExecRawQuery(arg1:string):Promise<models.SQLResult>;

export function GetTables():Promise<Array<string>>;

export function GetViews():Promise<Array<string>>;

export function NewDatabase(arg1:string):Promise<void>;

export function OpenDatabase():Promise<string>;

export function OpenInMemoryDatabase():Promise<string>;

export function SaveFile(arg1:string,arg2:string):Promise<boolean>;

export function SelectTable(arg1:string):Promise<models.SQLResult>;
