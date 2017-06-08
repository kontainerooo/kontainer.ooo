import { Component, OnInit, OnDestroy } from '@angular/core';
import { KmiTemplate } from '../../kmi-template';
import { GlobalDataService } from '../../../../services/global-data.service';

const template = {
	"name": "status",
	"parameters": [
		{
			"name": "start",
			"type": "emit"
		},
		{
			"name": "stop",
			"type": "emit"
		},
		{
			"name": "status", 
			"type": "poll",
			"interval": "500"
		},
		{
			"name": "title",
			"type": "value"
		} 
	]
};

@Component({
  selector: 'kro-kmi-status',
  templateUrl: './kmi-status.component.html',
  styleUrls: ['./kmi-status.component.scss']
})
export class KmiStatusComponent implements OnInit {
  title: string;
  status: number;
  statusInterval: NodeJS.Timer;

  constructor(private gds: GlobalDataService) {
    this.title = this.gds.getValueSnapshot(template.name, 'title');
    // this.statusInterval = setInterval(() => {
    //   this.gds.sendCommand('status').subscribe(
    //     value => {
    //       let intValue = parseInt(value);
    //       if(intValue == 1) {
    //         this.status = 2;
    //       } else if (intValue == 0) {
    //         this.status = 0;
    //       } else {
    //         this.status = 1;
    //       }
    //     },
    //     error => {
    //       console.log(error);
    //     }
    //   );
    // }, 500);
  }

  ngOnInit() {
    this.gds.sendCommand('status').subscribe(
      value => {
        let intValue = parseInt(value);
        if(intValue == 1) {
          this.status = 2;
        } else if (intValue == 0) {
          this.status = 0;
        } else {
          this.status = 1;
        }
      },
      error => {
        console.log(error);
      }
    );
  }

  ngOnDestroy() {
    // clearInterval(this.statusInterval);
  }

  toggleStatus() {
    if(this.status == 0) {
      this.gds.sendCommand('start').subscribe(
        value => {
          console.log(value);
          this.ngOnInit();
        },
        error => {
          console.log(error);
        }
      );
      this.status = 1;
    } else {
      this.gds.sendCommand('stop').subscribe(
        value => {
          console.log(value);
          this.ngOnInit();
        },
        error => {
          console.log(error);
        }
      );

      this.status = 1;
    }
  }

}
