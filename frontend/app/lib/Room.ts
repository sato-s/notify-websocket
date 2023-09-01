import * as neffos from "neffos.js";

export default class Room {
  url: string;

  constructor(url: string) {
    this.url = url;
    this.conn = await neffos.dial(url, {
    });
  }

  _OnNamespaceConnected(nsConn, msg) {
  }
  _OnNamespaceDisconnect(nsConn, msg) {
  }

}
