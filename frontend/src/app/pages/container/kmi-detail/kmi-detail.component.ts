import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs/Rx';

@Component({
  selector: 'kro-kmi-detail',
  templateUrl: './kmi-detail.component.html',
  styleUrls: ['./kmi-detail.component.scss']
})
export class KmiDetailComponent implements OnInit, OnDestroy {
  private routeId: number;
  private routeSub: Subscription;

  constructor(private route: ActivatedRoute) {
    console.log(this.route);
  }

  ngOnInit() {
    this.routeSub = this.route.params.subscribe(params => {
      this.routeId = params['id'];
    });
  }

  ngOnDestroy() {
    this.routeSub.unsubscribe();
  }
}
