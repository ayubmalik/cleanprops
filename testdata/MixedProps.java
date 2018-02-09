class Example3 {

    private String msg1;
    private String msg2;

    public Example3(@Value("#{cp.key3}") String key1, @Value("${cp.key4}") String key2) {
        this.msg1 = msg1;
        this.msg2 = msg2;
    }

    public void format() {
        return "Msg is: {}" + msg1 + msg2
    }

}
