import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs/Rx';
import { GlobalDataService } from '../../../services/global-data.service';
import { kmi } from '../../../../messages/messages';

@Component({
  selector: 'kro-kmi-detail',
  templateUrl: './kmi-detail.component.html',
  styleUrls: ['./kmi-detail.component.scss']
})
export class KmiDetailComponent implements OnInit, OnDestroy {
  private containerId: string;
  private refId: number;
  private containerName: string;
  private kmi: kmi.KMI;
  private routeSub: Subscription;

  private loaded: boolean;

  constructor(private route: ActivatedRoute, private gds: GlobalDataService) {
    this.loaded = false;
  }

  ngOnInit() {
    this.routeSub = this.route.params.subscribe(params => {
      this.containerId = params['id'];
      this.refId = params['refId'];
      this.containerName = params['name'];

      this.gds.setAndGetContainerKMI(this.containerName, this.containerId).subscribe(
        value => {
          this.kmi = value;
          this.loaded = true;
        },
        error => {
          console.log(error);
        }
      )
    });
  }

  hasTemplate(name: string): boolean {
    for(let module of this.kmi.frontend) {
      if(module.template == name) {
        return true;
      }
    }
    return false;
  }

  ngOnDestroy() {
    this.routeSub.unsubscribe();
  }
}
