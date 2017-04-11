import { Pipe, PipeTransform } from '@angular/core';
import { kmi } from '../../messages/messages';

@Pipe({
  name: 'searchKmi'
})
export class SearchKmiPipe implements PipeTransform {

  transform(value: Array<kmi.KMDI>, searchTerm: string): Array<kmi.KMDI> {
    if(searchTerm == undefined || searchTerm == '') {
      return value;
    }
    return value.filter((kmdi) => {
      return kmdi.name.toLocaleLowerCase().indexOf(searchTerm.toLocaleLowerCase()) > -1;
    });
  }

}
