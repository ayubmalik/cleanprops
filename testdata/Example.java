class Example {

    private String msg1;

    private String msg2;

    public Example(@Value("${cp.key1}") String key1, @Value("${cp.key2}") String key2) {
        this.msg1 = msg1;
        this.msg2 = msg2;
    }

    public void format() {
        return "Msg is: {}" + msg1 + msg2
    }

    @java.lang.Override
    public java.lang.String toString() {
        return "Example{" +
                "msg1='" + msg1 + '\'' +
                ", msg2='" + msg2 + '\'' +
                '}';
    }
}
