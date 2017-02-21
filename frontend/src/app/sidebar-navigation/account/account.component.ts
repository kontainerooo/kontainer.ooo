import { Component, Input, OnChanges, ElementRef, Renderer } from '@angular/core';

@Component({
  selector: 'kmi-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.scss']
})
export class AccountComponent implements OnChanges {
  @Input('bg') backgroundImage: string;
  @Input() realname: string;
  @Input() username: string;

  constructor(private elRef: ElementRef, private renderer: Renderer) { }

  ngOnChanges() {
    this.setBackground();
  }

  setBackground() {
    this.renderer.setElementStyle(this.elRef.nativeElement, 'background-image', `url('${this.backgroundImage}')`);
  }

}
