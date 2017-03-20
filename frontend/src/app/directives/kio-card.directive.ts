import { Directive } from '@angular/core';

@Directive({
  selector: '[kio-card-full-width]',
  host: {
    '[class.kio-card-full-width]': 'true'
  }
})
export class KioCardFullWidthDirective {

  constructor() { }

}
