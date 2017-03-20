import { kroPage } from './app.po';

describe('kontainer.ooo App', function() {
  let page: kroPage;

  beforeEach(() => {
    page = new kroPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
