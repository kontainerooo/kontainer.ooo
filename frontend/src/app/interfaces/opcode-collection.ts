export interface OpcodeCollection {
  [k: string]: {
    id: string,
    methods: {
      [k: string]: string
    }
  }
}