/* tslint:disable:no-unused-variable */

import { TestBed, async } from '@angular/core/testing';
import { SearchKmiPipe } from './search-kmi.pipe';
import { kmi } from '../../messages/messages';

describe('SearchKmiPipe', () => {
  it('create an instance', () => {
    const pipe = new SearchKmiPipe();
    expect(pipe).toBeTruthy();
  });

  it('should find the right kmis', () => {
    const pipe = new SearchKmiPipe();
    const kmiModules = [
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

    let value = pipe.transform(kmiModules, 'nod');
    expect(value.length).toEqual(1);
    for(let el of value) {
      expect(el.ID).toEqual(0);
    }

    value = pipe.transform(kmiModules, '.j');
    expect(value.length).toEqual(1);
    for(let el of value) {
      expect(el.ID).toEqual(0);
    }

    value = pipe.transform(kmiModules, 'n');
    expect(value.length).toEqual(2);
    for(let el in value) {
      expect(value[el].ID).toEqual(parseInt(el));
    }
  });
});
