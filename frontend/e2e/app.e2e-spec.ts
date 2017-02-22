import { Kontainer.IoPage } from './app.po';

describe('kontainer.io App', function() {
  let page: Kontainer.IoPage;

  beforeEach(() => {
    page = new Kontainer.IoPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
