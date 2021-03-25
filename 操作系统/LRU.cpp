//智能指针实现
class LRUCache {
private:
	struct node {
		int key, val;
		shared_ptr<node> next;
		weak_ptr<node> pre;
		node() : key(0), val(0) {}
		node(int _key, int _val) : key(_key), val(_val) {}
	};
	unordered_map<int, shared_ptr<node>> cache;
	shared_ptr<node> head;
	shared_ptr<node> tail;
	int size;
	int capacity;
public:
	LRUCache(int _capacity) : capacity(_capacity), size(0) {
		head = make_shared<node>();
		tail = make_shared<node>();
		head->next = tail;
		tail->pre = head;
	}

	int get(int key) {
		if (!cache.count(key)) {
			return -1;
		}

		auto pnode = cache[key];
		moveToHead(pnode);
		return pnode->val;
	}


	void put(int key, int value) {
		if (!cache.count(key)) {
			shared_ptr<node> pnode = make_shared<node>(key, value);
			cache[key] = pnode;
			addToHead(pnode);
			size++;
			if (size > capacity) {
				auto removed = removeTail();
				cache.erase(removed->key);
				size--;
			}
		} else {
			shared_ptr<node> pnode = cache[key];
			pnode->val = value;
			moveToHead(pnode);
		}
	}

	void addToHead(shared_ptr<node>& pnode) {
		pnode->pre = head;
		pnode->next = head->next;
		head->next->pre = pnode;
		head->next = pnode;
	}

	void removeNode(shared_ptr<node>& pnode) {
		pnode->pre.lock()->next = pnode->next;
		pnode->next->pre = pnode->pre;
	}

	void moveToHead(shared_ptr<node>& pnode) {
		removeNode(pnode);
		addToHead(pnode);
	}

	shared_ptr<node> removeTail() {
		shared_ptr<node> pnode = tail->pre.lock();
		removeNode(pnode);
		return pnode;
	}
};