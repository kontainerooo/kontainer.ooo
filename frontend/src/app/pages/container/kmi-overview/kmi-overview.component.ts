import { Component, OnInit } from '@angular/core';
import { kmi } from '../../../../messages/messages';

@Component({
  selector: 'kro-kmi-overview',
  templateUrl: './kmi-overview.component.html',
  styleUrls: ['./kmi-overview.component.scss']
})
export class KmiOverviewComponent {
  kmiModules: Array<kmi.KMDI>;
  kmiStatus: Array<boolean>;

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

    this.kmiStatus = [
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
