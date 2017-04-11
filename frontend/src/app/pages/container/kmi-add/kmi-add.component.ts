import { Component, OnInit } from '@angular/core';
import { kmi } from '../../../../messages/messages';

@Component({
  selector: 'kro-kmi-add',
  templateUrl: './kmi-add.component.html',
  styleUrls: ['./kmi-add.component.scss']
})
export class KmiAddComponent {
  kmiModules: Array<kmi.KMDI>;
  kmiInstalled: Array<boolean>;

  constructor() {
    this.kmiModules = [
      new kmi.KMDI({
        ID: 0,
        description: 'Event-driven I/O server-side JavaScript environment based on V8. Includes API documentation, change-log, examples and announcements.',
        name: 'Node.js',
        type: kmi.Type.WEBSERVER,
        version: '1.0.0+7.8.0'
      }),
      new kmi.KMDI({
        ID: 1,
        description: 'Building on the Best of Relational with the Innovations of NoSQL',
        name: 'MongoDB',
        type: kmi.Type.WEBSERVER,
        version: '1.0.0+3.4.3'
      })
    ];

    this.kmiInstalled = [
      true,
      false
    ];
  }

  convertType(type: kmi.Type): string {
    switch(type) {
      case 0:
        return 'Webserver';
      default:
        return 'Default';
    }
  }

}
