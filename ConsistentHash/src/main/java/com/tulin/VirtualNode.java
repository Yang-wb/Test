package com.tulin;

/**
 * Created by IntelliJ IDEA
 * User: Ginger
 * Date: 2018/8/17
 * Time: 9:14 AM
 */
public class VirtualNode<T extends Node> implements Node {
    private final T physicalNode;
    private final int replicaIndex;

    public VirtualNode(T physicalNode, int replicaIndex) {
        this.replicaIndex = replicaIndex;
        this.physicalNode = physicalNode;
    }

    @Override
    public String getKey() {
        return physicalNode.getKey() + "-" + replicaIndex;
    }

    public boolean isVirtualNodeOf(T pNode) {
        return physicalNode.getKey().equals(pNode.getKey());
    }

    public T getPhysicalNode() {
        return physicalNode;
    }
}