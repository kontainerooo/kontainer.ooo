import { Component, OnInit } from '@angular/core';
import { KmiTemplate } from '../../kmi-template';
// TODO example import as there is no service for this right now
import { KmiService } from '../../../../services/kmi.service';
import { kmi } from '../../../../../messages/messages';

@Component({
  selector: 'kro-kmi-status',
  templateUrl: './kmi-status.component.html',
  styleUrls: ['./kmi-status.component.scss']
})
export class KmiStatusComponent extends KmiTemplate implements OnInit {
  title: string;
  status: number;

  constructor(private kmiService: KmiService) {
    super(kmiService);

    this.title = this.getParameter('title');
    this.status = parseInt(this.getParameter('status'));

    // TODO remove mock data when service exists
    this.status = 0;
  }

  toggleStatus() {
    // TODO wait for service
    if(this.status == 0) {
      this.executeCommand('start');

      this.status = 1;
      setTimeout(() => {
        this.status = 2;
      }, 1000);
    } else {
      this.executeCommand('stop');

      this.status = 0;
    }
  }

  ngOnInit() {
    this.kmiService.messages.subscribe(
      (value) => {
        console.log(value);
      },
      (error) => {
        console.log(error);
      },
      () => {
        console.log('yup');
      }
    );

    this.kmiService.next('GetKMIRequest', {
      ID: 1
    });
  }

}
